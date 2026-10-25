package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 int16 = -1
    v4 := make(map[int]bool)
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{fi5}
        var v7 float64 = (1e10 + -2.25)
        fmt.Println(v7, v3)
        _ = fi5
        _ = v6
        _ = v7
    }
    v3 += -274
    var v8 uint16 = 32767
    fmt.Println(v8, v3)
    switch 1 {
    case 2147483647:
        fmt.Println(v3, v8)
    default:
        v8 -= (^(v8))
    }
    switch "hello" {
    case "a":
        switch "hello" {
        case "a":
            switch 9223372036854775807 {
            case 127:
                for fi9 := 0; fi9 < 19; fi9++ {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        var v11 int64 = -1
                        var v12 uint8 = 1
                        v13 := GenS1{F0: (v3 - 1), F1: -1, F2: 1}
                        v14 := []int{}
                        fmt.Println(v11, v11, fi10)
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi9
                }
            case -461:
                v8 += v8
            case 9223372036854775807:
                var v15 uint64 = 0
                _ = v15
            }
        default:
            v16 := GenS2{F0: 42, F1: 42}
            _ = v16
        }
    case "\n\t":
        v17 := []int{2147483647, -9223372036854775808, -628, -323}
        _ = v17
    case "":
        var v18 uint8 = 44
        _ = v18
    default:
        var v19 uint = 529
        _ = v19
    }
    _ = len(v4)
    var v20 uint = 4294967295
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := v4[-1]
        v20 += (v20 ^ v20)
        var v23 int = 127
        _ = fi21
        _ = v22
        _ = v23
    }
    if (1 <= 197) {
        var v24 uint16 = v8
        _ = v24
    } else {
        v25 := GenS2{F0: 42, F1: 2}
        _ = v25
    }
    v26 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v8
}
