package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int8; F1 int32; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 1
    v3 := make(map[int]uint32)
    if false {
        var v4 bool = false
        _ = v4
    }
    for fi5 := 0; fi5 < 13; fi5++ {
        v2 += v2
        for fi6 := 0; fi6 < 12; fi6++ {
            if (false && false) {
                fmt.Println(v2)
            }
            _ = fi6
        }
        var v7 uint64 = 240
        _ = fi5
        _ = v7
    }
    v8 := []int{259, 169}
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            v2 += v2
            var v11 uint32 = 4294967295
            _ = fi10
            _ = v11
        }
        var v12 uint8 = 0
        v12 -= (v12 ^ v12)
        _ = fi9
        _ = v12
    }
    if ("" <= "") {
        var v13 uint = 4294967295
        _ = v13
    } else {
        v14 := GenS1{F0: 4294967295, F1: (!(!(!true))), F2: v2}
        var v15 uint64 = 2
        v16 := GenS1{F0: 465, F1: true, F2: v2}
        _ = v14
        _ = v15
        _ = v16
    }
    if true {
        var v17 float64 = 0.0
        var v18 int32 = 2147483647
        if true {
            v2 += v2
            v19 := GenS1{F0: 42, F1: false, F2: (v2 & (v18 / ((v2) | 1)))}
            if true {
                _ = len(v1)
            }
            _ = v19
        }
        _ = v17
        _ = v18
    } else {
        fmt.Println(v2)
        if (1 == 128) {
            var v20 uint = 33
            _ = v20
        } else {
            v2 = v2
            fmt.Println(v2)
        }
        var v21 uint16 = 197
        _ = v21
    }
    v2 += (v2 * (v2 / (((^(v2))) | 1)))
    v22 := GenS2{F0: 255, F1: -1}
    v2 -= (v2 * -1)
    v23 := GenS3{F0: 1, F1: v2, F2: 32767}
    if true {
        _ = len(v3)
        var v24 int16 = 0
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v8
}
