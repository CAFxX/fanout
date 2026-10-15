package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 float32 = -2.25
    var v2 int64 = 42
    var v3 int32 = -835
    var v4 int = 32768
    if true {
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 uint32 = 65535
                v8 := []int{(fi5 >> ((0) & 63))}
                _ = fi6
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: 328, F1: 32767}
            v2 = v2
            _ = fi5
            _ = v9
        }
        v10 := GenS1{F0: 4294967295, F1: 65535}
        v11 := []int{}
        _ = v10
        _ = v11
    } else {
        v2 -= (v2 * (v2 % ((42) | 1)))
        v1 = 0.1
    }
    v12 := []int{v4, (v4 | (v4 / ((int(v3)) | 1))), (^(v4)), (v4 ^ v4)}
    var v13 uint64 = 42
    fmt.Println(v2)
    if (42 == -1) {
        var v14 uint64 = (v13 / (((v13 / ((v13) | 1))) | 1))
        _ = v14
    }
    v15 := GenS2{F0: (v3 % (((v3 & (v3 * 61))) | 1))}
    v16 := GenS1{F0: 2, F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
