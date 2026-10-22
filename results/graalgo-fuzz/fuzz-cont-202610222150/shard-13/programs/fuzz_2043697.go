package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    var v1 string = ("\n\t" + ("\n\t" + "\n\t"))
    v2 := make(map[string]uint32)
    var v3 int16 = -1
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    v6 := make(map[uint32]string)
    v7 := []int{4294967295}
    var v8 bool = true
    fmt.Println(v8, v3)
    fmt.Println(v3)
    var v9 uint16 = 18
    for fi10 := 0; fi10 < 9; fi10++ {
        v9 += v9
        _ = fi10
    }
    if v8 {
        for fi11 := 0; fi11 < 9; fi11++ {
            v7 = append(v7, (fi11 - 127))
            v1 = v1
            _ = fi11
        }
        var v12 float64 = (1e10 / 3.14159)
        _ = v12
    }
    v9 += 951
    v3 = 1
    v4[v1] = ((v1 + v1) == "foo bar")
    var v13 bool = (!(!(false || false)))
    v14 := GenS1{F0: 299, F1: 653}
    var v15 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
