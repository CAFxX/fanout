package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 bool }
func genFn0(p0 string, p1 int8) int64 {
    var v6 int64 = 42
    p0 = "a"
    _ = p0
    _ = p1
    _ = v6
    return (v6 << ((0) & 63))
}
func genFn1() int32 {
    v7 := []int{1, -1, -72}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint32 = 2147483647
        _ = fi8
        _ = v9
    }
    var v10 string = (("\n\t" + (("" + ("hello" + (("\n\t" + "hello") + "hello"))) + "hello")) + "\n\t")
    _ = v10
    _ = v7
    return -452
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 65535
    var v3 int8 = 127
    var v4 uint32 = (v2 ^ (^(v2)))
    v5 := make(map[int]uint32)
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := []int{2147483647, (fi11 % (((fi11 | (fi11 ^ fi11))) | 1)), fi11, 2147483647}
        _ = len(v12)
        _ = fi11
        _ = v12
    }
    v13 := []int{2147483647, 128}
    var v14 int64 = -9223372036854775808
    if false {
        fmt.Println(v3)
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS1{F0: -1, F1: 120, F2: (0 >= -1)}
        v4 -= v4
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint32 = v4
        fmt.Println(v2, v18)
        fmt.Println(fi17, v2)
        _ = fi17
        _ = v18
    }
    switch 127 {
    case -1:
        var v19 string = "foo bar"
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            _ = len(v13)
            fmt.Println(v2, fi20)
            _ = fi21
        }
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
