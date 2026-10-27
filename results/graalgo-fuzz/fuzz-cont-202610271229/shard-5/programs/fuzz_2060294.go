package main

import (
    "fmt"
)

func genFn0() int {
    var v4 string = ("a" + "hello")
    var v5 bool = (-2147483648 <= -723)
    var v6 float64 = (((1.5 + 3.14159) * 1.5) * float64((1e10 + -2.25)))
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := []uint32{247, 810, 1, 740}
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    var v7 uint64 = 42
    switch 4294967295 {
    case 127:
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int8 = 1
            fmt.Println(v7, v9, fi8)
            _ = fi8
            _ = v9
        }
    }
    v7 = v7
    var v10 uint16 = 765
    var v11 int = genFn0()
    v11 += -809
    switch "a" {
    case "":
        v7 += (v7 & 42)
    case "foo bar":
        if true {
            v11 += genFn0()
        }
    case "hello":
        fmt.Println(v11, v7, v11)
    }
    switch "" {
    case "a":
        var v12 int32 = 0
        _ = v12
    default:
        fmt.Println(v10)
    }
    fmt.Println(v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
}
