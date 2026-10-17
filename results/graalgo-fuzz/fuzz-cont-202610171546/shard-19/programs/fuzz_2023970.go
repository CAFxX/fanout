package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int8; F2 uint }

func main() {
    var v1 uint16 = 255
    var v2 int = -14
    v3 := make(map[string]int)
    if true {
        for fi4 := 0; fi4 < 17; fi4++ {
            var v5 int32 = -149
            _ = fi4
            _ = v5
        }
        fmt.Println(v2, v1)
    }
    var v6 uint64 = 4294967295
    v7 := []int{v2}
    var v8 int32 = 1
    _ = len(v3)
    fmt.Println(v2)
    v8 -= v8
    v2 = (-(v2))
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int32 = (v8 * v8)
        for fi11 := 0; fi11 < 15; fi11++ {
            if false {
                if true {
                    v12 := GenS1{F0: -173}
                    v1 += v1
                    _ = v12
                }
                fmt.Println(v10, v2)
                v13 := GenS2{F0: 4294967295, F1: 1, F2: 42}
                var v14 int16 = 32767
                if len(v7) > 0 {
                    v7[0] = v2
                }
                v15 := GenS2{F0: 255, F1: 100, F2: 911}
                if len(v7) > 0 {
                    v7[0] = (fi11 & 9223372036854775807)
                }
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 int = 626
            v17 := GenS1{F0: 708}
            _ = fi11
            _ = v16
            _ = v17
        }
        var v18 float32 = -2.25
        _ = fi9
        _ = v10
        _ = v18
    }
    switch v2 {
    case 32767:
        fmt.Println(v8)
    default:
        if len(v7) > 0 {
            v7[0] = v2
        }
    }
    v19 := GenS1{F0: 773}
    fmt.Println(v2, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
