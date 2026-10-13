package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16) int8 {
    v5 := []int{-706}
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    _ = v5
    return -60
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 2
    var v3 string = ("\n\t" + "foo bar")
    var v4 uint64 = 2
    var v6 uint = 1
    v3 = "hello"
    v7 := []int{32767, -9223372036854775808, -702}
    fmt.Println(v6, v2, v2)
    var v8 bool = (1 != 4294967295)
    switch v3 {
    case "a":
        var v9 float64 = ((0.1 / 0.1) + 0.0)
        _ = v9
    case "foo bar":
        v7 = append(v7, int(v2))
    }
    var v10 uint8 = v2
    v6 -= v6
    var v11 int8 = 100
    for fi12 := 0; fi12 < 16; fi12++ {
        v11 = (v11 & 52)
        _ = fi12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: -68}
        v4 -= 1
        _ = fi13
        _ = v14
    }
    v11 -= v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
