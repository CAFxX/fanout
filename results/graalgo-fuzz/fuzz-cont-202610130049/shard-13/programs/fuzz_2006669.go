package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint = 0
    var v2 int32 = 1
    var v3 uint32 = 613
    var v4 int = 600
    v5 := GenS1{F0: v3}
    if false {
        fmt.Println(v1, v4, v4)
    } else {
        var v6 int64 = -1
        fmt.Println(v3, v3, v1)
        _ = v6
    }
    var v7 uint16 = 32767
    if (1e10 != 3.14159) {
        v8 := []int{}
        for fi9 := 0; fi9 < 5; fi9++ {
            v2 += (v2 + v2)
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 int8 = -1
                fmt.Println(v4, v1, fi10)
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        var v12 uint32 = (v3 % ((v3) | 1))
        _ = v12
        _ = v8
    }
    fmt.Println(v4, v4, v7)
    v7 = (^(v7))
    v7 = (v7 / ((84) | 1))
    var v13 uint16 = 239
    v14 := GenS2{F0: 0, F1: -423}
    v15 := []int{32768, 0, (v4 << (((v1 << ((v1) & 63))) & 63)), -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
