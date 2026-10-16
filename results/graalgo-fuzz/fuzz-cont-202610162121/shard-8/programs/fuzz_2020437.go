package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []int{-1, 0, -586}
    v2 := make(map[int]bool)
    v3 := []uint32{1, 42, 318, 0}
    var v4 uint64 = 1
    v5 := []uint32{}
    v6 := GenS1{F0: 128}
    v4 -= v4
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    var v8 float32 = ((float32(1e10) + 1.5) - (0.1 / 1.5))
    v4 -= v4
    var v9 uint32 = 2147483648
    for fi10 := 0; fi10 < 13; fi10++ {
        delete(v2, (fi10 * (fi10 + -1)))
        _ = fi10
    }
    v4 = v4
    fmt.Println(v9, v9)
    if (true || false) {
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(fi11, fi11)
            _ = fi11
        }
    } else {
        if true {
            if false {
                var v12 bool = (false || (1 >= 42))
                _ = v12
            } else {
                v9 = v9
                _ = len(v2)
            }
            v9 -= (v9 % ((uint32(v4)) | 1))
        }
    }
    if (128 == 127) {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint8 = 1
            v14 = (v14 % ((192) | 1))
            v14 += v14
            _ = fi13
            _ = v14
        }
        v9 -= 96
    }
    var v15 bool = (float64(0.0) == 1.5)
    var v16 uint = 18446744073709551615
    if v15 {
        v5 = append(v5, v9)
        var v17 int8 = 100
        _ = v17
    }
    fmt.Println(v16, v16, v15)
    fmt.Println(v15, v15, v16)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
