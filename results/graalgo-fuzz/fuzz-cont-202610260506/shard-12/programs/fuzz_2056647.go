package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int64 }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]int)
    var v3 int = -814
    var v4 uint32 = 2147483648
    v5 := []int{(v3 - (v3 - v3))}
    var v6 int32 = 2147483647
    v7 := v2[-916]
    var v8 float64 = 1e10
    v6 += 993
    fmt.Println(v6, v8)
    fmt.Println(v3)
    var v9 int64 = 1
    if true {
        v8 = v8
        var v10 float32 = 0.1
        _ = v10
    }
    if false {
        fmt.Println(v3, v6)
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(v7, v9, fi11)
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: "hello", F1: 0, F2: int64(v6)}
                var v14 uint32 = (v4 + (v4 >> ((42) & 63)))
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 int16 = 1
            _ = fi11
            _ = v15
        }
    }
    v6 = v6
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v9 = (v9 << ((314) & 63))
        var v16 int16 = -1
        _ = v16
    }
    v4 += v4
    if false {
        _ = len(v5)
        v6 -= (v6 - v6)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
