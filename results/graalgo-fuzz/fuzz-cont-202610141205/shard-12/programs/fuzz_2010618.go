package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
func genFn0(p0 uint64, p1 bool) int {
    v5 := []int{-1}
    if p1 {
        p0 = 301
        if false {
            for fi6 := 0; fi6 < 16; fi6++ {
                _ = len(v5)
                var v7 int = fi6
                v7 -= (fi6 - 128)
                _ = fi6
                _ = v7
            }
        } else {
            v8 := GenS1{F0: 1, F1: (p0 * 44), F2: true}
            fmt.Println(p1)
            _ = v8
        }
    }
    if (-875 != 1) {
        for fi9 := 0; fi9 < 13; fi9++ {
            fmt.Println(p0)
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return -1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1.5
    var v3 uint = 0
    var v4 float32 = v2
    v3 = v3
    v2 = v2
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: 2, F1: 2, F2: false}
        _ = fi10
        _ = v11
    }
    switch "a" {
    case "\n\t":
        v3 = v3
    default:
        v1[1] = (!false)
    }
    var v12 uint64 = 18446744073709551615
    delete(v1, 152)
    _ = len(v1)
    if false {
        if false {
            v2 = 1e10
            for fi13 := 0; fi13 < 19; fi13++ {
                _ = len(v1)
                v4 = 0.0
                var v14 int32 = 650
                _ = fi13
                _ = v14
            }
            var v15 float32 = v2
            _ = v15
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
