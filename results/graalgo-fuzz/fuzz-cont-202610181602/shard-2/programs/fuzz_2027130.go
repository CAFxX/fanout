package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    v3 := []int{42, -799}
    var v4 bool = false
    delete(v1, -1)
    if len(v2) > 0 {
        v2[0] = 32768
    }
    v4 = (-1 >= 9223372036854775807)
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(v4, fi5)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int64 = 0
        if true {
            for fi8 := 0; fi8 < 14; fi8++ {
                v7 -= v7
                _ = fi8
            }
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int = 9223372036854775807
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 586, F1: 410}
    v12 := GenS2{F0: 1, F1: 2, F2: 248}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
