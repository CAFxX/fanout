package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint32 }

func main() {
    var v1 uint32 = 67
    var v2 uint16 = 255
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 bool = (!(true && false))
    v6 := make(map[int]int)
    fmt.Println(v2, v1, v1)
    v7 := []int{-167}
    var v8 uint8 = 42
    fmt.Println(v8, v5)
    fmt.Println(v1, v5, v2)
    v9 := GenS1{F0: (^(v8)), F1: (v1 & 65535), F2: 602}
    v2 += v2
    if false {
        v8 += v8
        fmt.Println(v2)
        if v5 {
            var v10 int8 = 0
            var v11 uint8 = (v8 % ((255) | 1))
            _ = v10
            _ = v11
        } else {
            v12 := GenS1{F0: v8, F1: v1, F2: (^(v1))}
            _ = v12
        }
    }
    var v13 int = -2147483648
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v14 := GenS1{F0: 1, F1: 4294967295, F2: (v1 + (v1 + v1))}
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            v13 = (fi15 | fi15)
            var v17 int64 = 0
            v7 = append(v7, (v13 * -2147483648))
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            v6[fi15] = (fi18 * (v13 & v13))
            var v19 int16 = 1000
            fmt.Println(v5, v2, v2)
            _ = fi18
            _ = v19
        }
        fmt.Println(fi15)
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
