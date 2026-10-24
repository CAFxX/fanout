package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    var v5 float64 = 1e10
    v6 := []int{-190, -1, 9223372036854775807}
    var v7 uint32 = (^(v1))
    var v8 int8 = 0
    v9 := v4[(v1 & 2147483647)]
    if true {
        v10 := v2[(v7 & 4294967295)]
        _ = v10
    }
    v11 := GenS1{F0: 436, F1: v7}
    var v12 uint64 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 uint16 = 65535
    if true {
        v14 := GenS2{F0: -203}
        switch "a" {
        case "\n\t":
            v1 += (v1 & (v1 ^ 833))
        }
        _ = v14
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 19; fi16++ {
                v8 += (v8 | (v8 + (v8 % (((v8 - (v8 % ((int8(v12)) | 1)))) | 1))))
                var v17 int64 = -149
                v13 = v13
                _ = fi16
                _ = v17
            }
            v7 += (v1 & (v1 + 65535))
            _ = fi15
        }
    }
    v18 := GenS2{F0: -114}
    if true {
        fmt.Println(v12)
        v7 -= (v7 | v1)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v19 int8 = (v8 & 80)
        _ = v19
    }
    fmt.Println(v5, v8, v12)
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
