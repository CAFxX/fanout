package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 string, p1 int16) int32 {
    var v5 int64 = 1
    fmt.Println(p1)
    for fi6 := 0; fi6 < 9; fi6++ {
        p1 = p1
        v7 := []int{fi6, 9223372036854775807, -528, fi6}
        _ = fi6
        _ = v7
    }
    if true {
        var v8 int64 = (v5 >> ((0) & 63))
        v9 := GenS1{F0: "foo bar", F1: -267}
        p1 -= p1
        _ = v8
        _ = v9
    }
    v5 -= (v5 % (((v5 ^ 9223372036854775807)) | 1))
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{}
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    fmt.Println(v1)
    var v10 uint32 = 4294967295
    if (true && true) {
        if false {
            v10 += v10
        } else {
            v11 := v4[""]
            _ = len(v2)
            _ = v11
        }
    }
    v12 := GenS2{F0: 620}
    var v13 float64 = -2.25
    var v14 uint16 = 65535
    for fi15 := 0; fi15 < 8; fi15++ {
        v13 = v13
        if false {
            _ = len(v2)
            v16 := GenS3{F0: 42}
            _ = v16
        }
        var v17 float64 = float64(float32(v13))
        _ = fi15
        _ = v17
    }
    if (true && false) {
        if (287 <= 2) {
            v18 := GenS1{F0: "", F1: -312}
            _ = v18
        } else {
            if false {
                v19 := GenS4{F0: 42, F1: 4294967295, F2: v14}
                _ = v19
            }
            v10 = 2
            var v20 float32 = float32(1.5)
            _ = v20
        }
        v21 := GenS1{F0: "", F1: v1}
        _ = v21
    }
    var v22 bool = (0 <= 42)
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(v10)
        _ = fi23
    }
    v24 := GenS5{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
