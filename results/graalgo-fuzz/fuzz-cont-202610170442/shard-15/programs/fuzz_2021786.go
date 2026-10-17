package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 bool }

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 74
    var v3 int32 = 42
    var v4 bool = (0.0 <= 0.0)
    fmt.Println(v4, v2)
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int8 = v2
        var v7 string = ("hello" + "a")
        if true {
            for fi8 := 0; fi8 < 16; fi8++ {
                fmt.Println(v7, v6)
                _ = fi8
            }
        } else {
            v9 := v1[-1]
            v10 := []int{fi5, fi5}
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v11 int = 549
    v12 := GenS1{F0: 4294967295, F1: (v3 / ((v3) | 1)), F2: v4}
    var v13 int16 = -812
    v4 = (32767 <= 2)
    fmt.Println(v2, v13)
    fmt.Println(v13)
    v14 := []int{}
    var v15 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
