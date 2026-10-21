package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 bool; F2 uint32 }

func main() {
    var v1 float32 = (1.5 / 1.5)
    v2 := []int{0, -590}
    var v3 uint8 = 13
    v4 := []uint32{uint32(v3), 966, 65535}
    v5 := make(map[uint32]string)
    var v6 uint16 = 255
    v7 := GenS1{F0: 1000, F1: 1}
    for fi8 := 0; fi8 < 4; fi8++ {
        switch fi8 {
        case -1:
            v6 = (v6 | (^(v6)))
        case -461:
            if (!false) {
                v6 = 368
                var v9 uint64 = 869
                for fi10 := 0; fi10 < 16; fi10++ {
                    v3 = (v3 | v3)
                    var v11 uint8 = (v3 | (v3 << ((4294967295) & 63)))
                    fmt.Println(v9)
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                var v12 int32 = 42
                v13 := GenS2{F0: (v12 | v12)}
                _ = v12
                _ = v13
            }
        case 9223372036854775807:
            v14 := GenS2{F0: 1}
            _ = v14
        }
        _ = fi8
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint16 = 543
        fmt.Println(v3, fi15, v6)
        _ = fi15
        _ = v16
    }
    fmt.Println(v6, v3, v3)
    v17 := GenS3{F0: -281, F1: true, F2: 2147483647}
    var v18 int32 = -9
    v19 := GenS2{F0: v18}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
