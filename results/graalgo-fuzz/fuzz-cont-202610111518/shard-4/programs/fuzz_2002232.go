package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    v1 := []int{}
    var v2 uint32 = 988
    v3 := make(map[string]int)
    var v4 int64 = -1
    v5 := []int{-1, -562, 2147483647, 32767}
    v6 := []uint32{(v2 - (v2 + (v2 & v2)))}
    if (-1 > 1) {
        v7 := GenS1{F0: 422}
        var v8 uint64 = 42
        v8 -= 2
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 32767}
    v10 := GenS1{F0: 32767}
    var v11 uint64 = 170
    v2 += v2
    for fi12 := 0; fi12 < 8; fi12++ {
        if true {
            if true {
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 int = fi13
                    v15 := GenS2{F0: ("\n\t" + ""), F1: 7}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            } else {
                v4 -= (v4 << ((42) & 63))
            }
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 uint = uint(v2)
                _ = fi16
                _ = v17
            }
        } else {
            v11 = v11
            fmt.Println(fi12)
        }
        v4 += (-(v4))
        fmt.Println(fi12, v2, fi12)
        _ = fi12
    }
    var v18 float64 = 1e10
    if len(v1) > 0 {
        v1[0] = -830
    }
    v2 = 2
    v18 -= float64(0.1)
    v11 = v11
    var v19 float32 = 1.5
    var v20 uint16 = 65535
    for fi21 := 0; fi21 < 10; fi21++ {
        if true {
            var v22 float32 = v19
            v23 := GenS1{F0: v20}
            fmt.Println(v19)
            _ = v22
            _ = v23
        } else {
            v20 = 65535
            v19 = (((v19 / (v19 - 0.1)) + ((1.5 - v19) - v19)) * v19)
            fmt.Println(v2, v11)
        }
        _ = fi21
    }
    v6 = append(v6, uint32(v11))
    fmt.Println(v19)
    var v24 int8 = -128
    fmt.Println(v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
