package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint32; F1 int64 }

func main() {
    var v1 int32 = -1
    var v2 uint8 = 90
    v3 := []int{4294967295, -1, -13}
    v4 := make(map[int]string)
    v5 := GenS1{F0: -1, F1: -32768}
    if len(v3) > 0 {
        v3[0] = 450
    }
    v2 += v2
    if true {
        fmt.Println(v1)
    } else {
        _ = len(v3)
        v1 = (v1 << ((4294967295) & 63))
    }
    _ = len(v3)
    var v6 uint = 1
    v7 := GenS2{F0: 712, F1: -1}
    v3 = append(v3, -1)
    v2 += v2
    if false {
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                v6 += (v6 % ((uint(v2)) | 1))
                _ = fi9
            }
            _ = fi8
        }
        v1 -= v1
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS3{F0: 0, F1: 9223372036854775807}
            var v12 int64 = 1
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v6 = 255
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int8 = 100
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
