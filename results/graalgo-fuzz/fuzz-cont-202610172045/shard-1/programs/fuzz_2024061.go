package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int64; F1 uint8 }

func main() {
    var v1 string = ""
    var v2 bool = true
    v3 := make(map[int]uint32)
    var v4 int64 = 42
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 string = v1
        _ = fi6
        _ = v7
    }
    v8 := v5[v1]
    v2 = true
    var v9 int8 = 0
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v2)
        v11 := []int{v8, fi10}
        _ = fi10
        _ = v11
    }
    var v12 int16 = 144
    if (v8 > (v8 | -102)) {
        v13 := GenS1{F0: v9}
        v14 := GenS2{F0: v4, F1: v2}
        _ = v13
        _ = v14
    }
    var v15 uint64 = 0
    v12 -= (v12 ^ (v12 * v12))
    fmt.Println(v9, v9)
    var v16 uint64 = (v15 * (v15 / ((v15) | 1)))
    if (!v2) {
        v17 := []int{}
        if true {
            if v2 {
                v17 = append(v17, 2147483647)
                for fi18 := 0; fi18 < 19; fi18++ {
                    v1 = v1
                    v16 = (^(v16))
                    _ = fi18
                }
            }
            var v19 int64 = v4
            for fi20 := 0; fi20 < 8; fi20++ {
                _ = len(v17)
                for fi21 := 0; fi21 < 13; fi21++ {
                    v12 = v12
                    var v22 int16 = (v12 * 393)
                    v23 := GenS3{F0: v19, F1: 127}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = fi20
            }
            _ = v19
        }
        v9 = (v9 | (v9 << ((4294967295) & 63)))
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
