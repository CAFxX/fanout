package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }

func main() {
    var v1 int64 = -273
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    var v4 int8 = 100
    v1 -= (v1 | (v1 / (((-(v1))) | 1)))
    var v5 float32 = (((1e10 * float32(float64(1e10))) - 1e10) - (3.14159 + 3.14159))
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            if (!false) {
                if (!true) {
                    v5 += v5
                    v8 := GenS1{F0: (fi6 - fi7), F1: (!(v2 >= (v2 + v2)))}
                    _ = v8
                } else {
                    v3[2147483647] = "foo bar"
                }
            }
            if true {
                var v9 uint = 31
                var v10 bool = false
                var v11 uint8 = 56
                var v12 int16 = 1000
                var v13 int64 = (v1 + int64(v11))
                _ = v10
                _ = v11
                _ = v12
                _ = v13
                _ = v9
            } else {
                var v14 float32 = 0.1
                v15 := []int{int(v4)}
                _ = len(v3)
                _ = v14
                _ = v15
            }
            _ = fi7
        }
        v16 := []int{(fi6 >> ((65535) & 63))}
        _ = fi6
        _ = v16
    }
    var v17 int32 = (^(v2))
    v18 := []int{}
    var v19 uint32 = 0
    v4 = (v4 << ((1) & 63))
    delete(v3, (^(v19)))
    var v20 int8 = int8(v2)
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 uint8 = 134
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 567, F1: (false || (!true))}
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 int32 = (^(v17))
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
