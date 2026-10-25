package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint8 = 105
    v2 := make(map[uint32]string)
    var v3 uint8 = v1
    fmt.Println(v3, v1)
    v4 := []int{-1, -9223372036854775808, 32768, 0}
    var v5 uint8 = v3
    if (-478 <= 1000) {
        if true {
            var v6 int8 = 0
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 float64 = (0.0 - 0.0)
                fmt.Println(v1, v6, v1)
                v9 := v2[42]
                _ = fi7
                _ = v8
                _ = v9
            }
            if true {
                var v10 int64 = 0
                var v11 int = -1
                var v12 uint = 42
                _ = v10
                _ = v11
                _ = v12
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := GenS1{F0: 1}
                    _ = fi13
                    _ = v14
                }
            }
            _ = v6
        }
        fmt.Println(v5)
        fmt.Println(v3, v1)
    }
    var v15 int64 = int64(v5)
    var v16 uint64 = 18446744073709551615
    v5 -= (v3 % (((v5 >> ((4294967295) & 63))) | 1))
    for fi17 := 0; fi17 < 15; fi17++ {
        delete(v2, 0)
        _ = fi17
    }
    fmt.Println(v15, v5, v1)
    for fi18 := 0; fi18 < 17; fi18++ {
        _ = len(v4)
        var v19 uint = 0
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: (v16 - v16), F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
