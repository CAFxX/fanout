package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 uint8 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]uint32)
    v3 := []int{129, 1}
    if true {
        v4 := GenS1{F0: 2147483647}
        v5 := GenS1{F0: 64}
        _ = v4
        _ = v5
    }
    var v6 int64 = -9223372036854775808
    v7 := GenS2{F0: 255, F1: 65535}
    fmt.Println(v1, v6)
    var v8 float64 = (float64((-2.25 + float32((1e10 - (float64((-2.25 - float32(1e10))) / float64(-2.25)))))) * -2.25)
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 2147483648}
        var v11 uint = 2
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 = v8
    v12 := GenS3{F0: v6, F1: (v1 - (v1 - 42))}
    v8 = v8
    v8 = v8
    if ((true && true) && true) {
        v6 = v6
    } else {
        fmt.Println(v6, v6)
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v13 int64 = (v6 | (v6 / (((v6 + (v6 / (((v6 >> ((42) & 63))) | 1)))) | 1)))
            var v14 int8 = -1
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 bool = false
                fmt.Println(v1, v16)
                _ = fi15
                _ = v16
            }
            v17 := GenS3{F0: 0, F1: v1}
            v6 = 1
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
