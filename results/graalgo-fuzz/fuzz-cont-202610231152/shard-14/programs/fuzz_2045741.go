package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64; F1 int; F2 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
type GenS4 struct { F0 int8; F1 uint8 }

func main() {
    var v1 uint = 101
    var v2 int16 = 0
    var v3 uint8 = 127
    v4 := []int{-858, 128}
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 bool = false
        _ = fi5
        _ = v6
    }
    switch "hello" {
    case "":
        var v7 uint32 = 2
        _ = v7
    }
    var v8 int32 = -315
    v8 -= v8
    switch -47 {
    case 128:
        v2 = v2
    case -1:
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 bool = true
            _ = fi9
            _ = v10
        }
    default:
        v11 := GenS1{F0: false, F1: 312}
        _ = v11
    }
    var v12 int8 = 1
    var v13 float64 = (0.1 - 0.1)
    v2 -= v2
    var v14 int64 = 0
    if (!(!true)) {
        v15 := GenS2{F0: (v14 | (v14 * int64(v3))), F1: 128, F2: int8(v2)}
        switch -631 {
        case -1:
            v16 := GenS3{F0: 4294967295, F1: 4294967295, F2: true}
            _ = v16
        case 128:
            v8 -= v8
        case 2147483647:
            v17 := GenS4{F0: v12, F1: v3}
            _ = v17
        }
        _ = v15
    }
    var v18 uint = v1
    _ = len(v4)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v4) > 0 {
            v4[0] = -867
        }
        var v19 uint = v1
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
