package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    v1 := []int{9223372036854775807, 195}
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v4 bool = true
    v5 := []int{-2147483648}
    v6 := make(map[int]string)
    var v7 uint8 = 255
    v8 := GenS1{F0: 210}
    v9 := GenS1{F0: 784}
    v7 = v7
    v10 := GenS1{F0: 1}
    var v11 int = -801
    v11 -= v11
    var v12 string = "foo bar"
    var v13 uint16 = 687
    if false {
        v14 := GenS2{F0: 42, F1: 42}
        for fi15 := 0; fi15 < 9; fi15++ {
            if v4 {
                if v4 {
                    v4 = (!(2147483647 > 492))
                }
                var v16 int32 = 42
                _ = v16
            }
            fmt.Println(v4)
            _ = fi15
        }
        fmt.Println(v11, v7)
        _ = v14
    } else {
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS2{F0: v11, F1: 120}
            v13 += (v13 & v13)
            v19 := GenS1{F0: -1}
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    var v20 float32 = ((-2.25 / (float32(-2.25) - (1.5 * 0.0))) / -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
