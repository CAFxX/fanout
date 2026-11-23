package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 724
    var v3 int16 = 32767
    v1[127] = "\n\t"
    v4 := GenS1{F0: 42, F1: "a", F2: 2147483647}
    var v5 uint32 = 1
    var v6 uint64 = 2
    var v7 float64 = 1.5
    fmt.Println(v2, v2)
    var v8 int32 = v2
    var v9 int8 = -62
    _ = len(v1)
    fmt.Println(v5, v3, v5)
    v10 := v1[-1]
    v11 := []int{0, 581, 32768}
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint = 856
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 int16 = -1
    v8 -= v2
    delete(v1, 0)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(v9, v15, v7)
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
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
