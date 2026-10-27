package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 string = "\n\t"
    var v6 int32 = -2147483648
    var v7 int64 = -1
    if ((1e10 * 1e10) == 1e10) {
        if true {
            v6 -= (v6 % (((v6 - 702)) | 1))
            if false {
                var v8 uint16 = 1
                v6 += (v6 << ((65535) & 63))
                v9 := v4[v5]
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 uint64 = 18446744073709551615
                    _ = fi10
                    _ = v11
                }
            }
            fmt.Println(v7, v5)
        } else {
            fmt.Println(v7, v7, v5)
        }
        v6 += v6
        delete(v4, v5)
    }
    v12 := GenS1{F0: 32767, F1: (-(v1))}
    v13 := GenS1{F0: -32768, F1: v1}
    v14 := []int{(-(v1)), (v1 * v1)}
    v1 = v1
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS2{F0: (v6 - v6), F1: false}
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 18446744073709551615
    var v18 uint8 = 145
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
