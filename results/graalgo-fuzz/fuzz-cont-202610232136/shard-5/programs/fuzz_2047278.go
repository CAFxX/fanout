package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint16 = 613
    v2 := make(map[uint32]uint32)
    var v3 float64 = float64(float32(0.1))
    v4 := make(map[string]bool)
    if true {
        v5 := GenS1{F0: -424, F1: -32768, F2: 705}
        v6 := GenS1{F0: 0, F1: 0, F2: 881}
        _ = v5
        _ = v6
    } else {
        if false {
            if true {
                v1 += 32767
            } else {
                switch 32768 {
                case -9223372036854775808:
                    v7 := GenS2{F0: 432}
                    _ = v7
                }
            }
            var v8 float64 = float64(3.14159)
            _ = v8
        }
        var v9 uint8 = 1
        _ = v9
    }
    var v10 uint32 = 2147483647
    v11 := []int{}
    if len(v11) > 0 {
        v11[0] = 128
    }
    if false {
        var v12 uint64 = 366
        _ = v12
    } else {
        for fi13 := 0; fi13 < 3; fi13++ {
            v1 += (^(v1))
            _ = fi13
        }
        v3 -= v3
        v14 := GenS3{F0: -128}
        _ = v14
    }
    _ = len(v11)
    v15 := GenS2{F0: 777}
    var v16 int64 = -9223372036854775808
    v1 += v1
    v16 -= 9223372036854775807
    v3 -= v3
    var v17 uint16 = (v1 & v1)
    for fi18 := 0; fi18 < 7; fi18++ {
        v10 = (v10 ^ (v10 / (((v10 | 65535)) | 1)))
        v11 = append(v11, 128)
        var v19 int8 = 127
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 443
    v21 := v2[(v10 ^ uint32(v1))]
    fmt.Println(v20, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
