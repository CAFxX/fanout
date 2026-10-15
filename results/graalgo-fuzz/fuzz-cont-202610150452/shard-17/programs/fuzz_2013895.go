package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    var v3 int64 = 0
    var v4 uint8 = 90
    fmt.Println(v2, v2)
    switch 916 {
    case 127:
        if false {
            fmt.Println(v3, v3, v2)
        }
    case 0:
        for fi5 := 0; fi5 < 4; fi5++ {
            v6 := GenS1{F0: 42}
            _ = fi5
            _ = v6
        }
    case 4294967295:
        var v7 uint16 = 255
        _ = v7
    }
    v8 := GenS1{F0: 0}
    switch -997 {
    case 232:
        _ = len(v1)
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := []int{fi9, (fi9 | (fi9 & (fi9 % (((fi9 | fi9)) | 1)))), (fi9 - fi9), (fi9 | 128)}
        _ = fi9
        _ = v10
    }
    if false {
        v11 := GenS1{F0: 582}
        var v12 uint8 = v4
        v13 := v1[1]
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 uint = 42
    v15 := v1[2]
    v2 += v2
    var v16 int64 = v3
    v17 := []int{(v15 ^ v15)}
    switch v15 {
    case 32767:
        v18 := GenS2{F0: 127, F1: v14}
        _ = v18
    default:
        var v19 int = (^(v15))
        _ = v19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v14 = 197
        delete(v1, uint32(v4))
        _ = fi20
    }
    fmt.Println(v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
