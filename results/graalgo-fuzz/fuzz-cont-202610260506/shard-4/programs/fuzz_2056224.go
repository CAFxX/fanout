package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int64 = -1
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    var v4 float32 = ((0.1 / 3.14159) - 1.5)
    var v5 uint64 = 295
    v1 = int64(v5)
    v6 := []int{120, -1, -2147483648, 487}
    var v7 float64 = 0.1
    v4 = v4
    if false {
        fmt.Println(v4, v4)
        v8 := GenS1{F0: (false && (true && false))}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v8
    }
    v4 = 1.5
    var v9 float64 = v7
    if true {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS2{F0: 1}
            v6 = append(v6, (fi10 >> ((uint(v5)) & 63)))
            _ = fi10
            _ = v11
        }
        v12 := GenS3{F0: v1}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v12
    }
    v7 += (v9 + (v9 + -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
