package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    v1 := []int{204}
    var v2 uint8 = 136
    var v3 int = -1
    var v4 int8 = 127
    v4 = (v4 - v4)
    if true {
        var v5 int64 = -433
        _ = v5
    } else {
        v2 = v2
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(fi6)
            var v7 int = (v3 * (v3 >> ((18446744073709551615) & 63)))
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint8 = v2
            v10 := GenS1{F0: 255, F1: v3}
            if (!false) {
                var v11 bool = true
                var v12 uint16 = 255
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v13 := GenS2{F0: v4, F1: 1}
    var v14 uint32 = 65535
    _ = len(v1)
    v3 = v3
    v2 = 173
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v14 = v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
