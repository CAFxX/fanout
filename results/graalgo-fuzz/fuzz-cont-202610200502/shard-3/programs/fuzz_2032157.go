package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int8) bool {
    switch "a" {
    case "foo bar":
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(fi6, fi6)
            v7 := []int{fi6}
            var v8 uint8 = 1
            _ = fi6
            _ = v7
            _ = v8
        }
    case "\n\t":
        v9 := []int{-464, 2147483647, -2147483648}
        _ = v9
    }
    if (-832 < 125) {
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int8 = (p0 % ((123) | 1))
                _ = fi11
                _ = v12
            }
            var v13 int = fi10
            v14 := GenS1{F0: 4294967295, F1: 65535}
            _ = fi10
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: 42, F1: 2}
        _ = v15
    }
    _ = p0
    return true
}

func main() {
    var v1 bool = false
    v2 := make(map[int]uint32)
    v3 := []int{610, -1, -895}
    var v4 uint16 = 42
    var v5 uint16 = (v4 >> ((2) & 63))
    var v16 uint = 650
    _ = len(v2)
    v5 = v5
    fmt.Println(v4, v4)
    v1 = v1
    fmt.Println(v5, v1, v1)
    v5 = (v5 & 2)
    v4 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
