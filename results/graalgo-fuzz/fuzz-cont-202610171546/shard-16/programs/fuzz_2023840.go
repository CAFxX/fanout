package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint16 = 344
    v2 := make(map[uint32]bool)
    var v3 int8 = -13
    v4 := []uint32{}
    var v5 uint16 = v1
    v6 := []int{32768, 42}
    fmt.Println(v5, v1)
    v7 := GenS1{F0: 2147483648}
    if false {
        var v8 bool = (998 <= 1)
        _ = v8
    }
    var v9 uint64 = 893
    v10 := GenS1{F0: 2147483647}
    var v11 uint32 = 4294967295
    var v12 int32 = -896
    if ((v9 & v9) == v9) {
        if false {
            if (!false) {
                var v13 bool = (-1 != -128)
                _ = v13
            } else {
                v14 := GenS1{F0: v11}
                fmt.Println(v1, v9, v1)
                for fi15 := 0; fi15 < 2; fi15++ {
                    v16 := GenS2{F0: fi15}
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
        } else {
            if false {
                v1 -= v5
                var v17 uint64 = v9
                _ = v17
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                v1 += v5
                _ = fi18
            }
            v19 := GenS2{F0: 849}
            _ = v19
        }
        var v20 int16 = 655
        for fi21 := 0; fi21 < 11; fi21++ {
            _ = len(v2)
            _ = fi21
        }
        _ = v20
    } else {
        var v22 int32 = (^(v12))
        _ = v22
    }
    v23 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
