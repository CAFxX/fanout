package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint64 = 19
    v2 := []int{-1, 128}
    v3 := []uint32{247, 102, 4294967295}
    fmt.Println(v1)
    v1 -= v1
    v1 += v1
    if true {
        fmt.Println(v1)
    } else {
        var v4 int16 = 32767
        _ = v4
    }
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint32 = 695
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi6
            _ = v7
        }
        v1 += 687
        if len(v3) > 0 {
            v3[0] = 919
        }
        _ = fi5
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 := GenS1{F0: 42, F1: 4294967295}
    if false {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS2{F0: -9223372036854775808}
            var v11 string = "\n\t"
            _ = fi9
            _ = v10
            _ = v11
        }
        v1 = 1
    }
    v1 = 0
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
