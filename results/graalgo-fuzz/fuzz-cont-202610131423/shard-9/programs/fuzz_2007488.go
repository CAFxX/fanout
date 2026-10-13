package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
func genFn0(p0 int8, p1 int) int {
    var v4 uint64 = 0
    v5 := GenS1{F0: 364, F1: "foo bar", F2: 1}
    if true {
        p0 += (p0 >> ((4294967295) & 63))
        var v6 bool = (42 != 413)
        fmt.Println(p1, p0)
        _ = v6
    } else {
        p1 += (p1 & p1)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    v1 := []int{753, 42, 32768}
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    v7 := v3["a"]
    switch "foo bar" {
    case "\n\t":
        v7 = v7
    case "":
        v7 = ((v7 + v7) + v7)
    }
    var v8 int16 = -1
    switch -1 {
    case -9223372036854775808:
        var v9 int32 = 1
        _ = v9
    case -378:
        v8 += (v8 - v8)
    case -568:
        var v10 string = "hello"
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(fi11, v8, v7)
        var v12 float64 = -2.25
        var v13 int = 1
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint16 = 32767
    v1 = append(v1, -150)
    if true {
        fmt.Println(v8)
        v14 += (v14 | (v14 & (v14 | v14)))
        var v15 int = -2147483648
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
