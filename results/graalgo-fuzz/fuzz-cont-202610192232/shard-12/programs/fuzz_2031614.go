package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []int{-968}
    var v2 int32 = -728
    v3 := make(map[uint32]string)
    var v4 uint32 = 42
    fmt.Println(v2)
    if true {
        var v5 uint32 = v4
        for fi6 := 0; fi6 < 19; fi6++ {
            v4 = v5
            _ = fi6
        }
        _ = v5
    }
    fmt.Println(v4, v2)
    v2 -= v2
    var v7 uint64 = 4294967295
    v8 := GenS1{F0: 20}
    var v9 uint = 2
    if ("hello" != "a") {
        if (128 <= -9223372036854775808) {
            var v10 bool = (!(-44 != 21))
            _ = v10
        }
    } else {
        v2 = v2
        if false {
            v11 := GenS1{F0: 0}
            if ((true && false) || (!false)) {
                var v12 uint32 = v4
                _ = v12
            } else {
                fmt.Println(v9, v7, v7)
                fmt.Println(v7, v9, v2)
            }
            v9 = (v9 % ((1) | 1))
            _ = v11
        } else {
            var v13 int64 = 381
            switch 986 {
            case -9223372036854775808:
                var v14 int = 9223372036854775807
                _ = v14
            }
            v15 := GenS2{F0: uint64(v4), F1: -522}
            _ = v13
            _ = v15
        }
        v7 = 251
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
