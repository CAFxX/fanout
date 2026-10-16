package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int8) int16 {
    v5 := []int{128, 48}
    var v6 int = 0
    var v7 int = -9223372036854775808
    var v8 int32 = 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 int = -275
    var v2 int16 = -1
    v3 := make(map[int]int)
    v4 := make(map[uint32]int)
    if (1e10 == 0.1) {
        v2 = -808
        v2 = (v2 | genFn0(-1))
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int64 = -607
        v1 = v1
        _ = fi9
        _ = v10
    }
    var v11 int = v1
    v12 := []int{558, 32767, (v11 | v1), (v11 % (((v11 << ((65535) & 63))) | 1))}
    var v13 int8 = 1
    var v14 uint8 = 255
    for fi15 := 0; fi15 < 15; fi15++ {
        if true {
            v2 -= 42
            v16 := GenS1{F0: v14}
            switch fi15 {
            case 229:
                var v17 uint32 = 681
                _ = v17
            case 127:
                v1 = v1
            default:
                if true {
                    fmt.Println(v1, fi15)
                } else {
                    _ = len(v4)
                    v3[v11] = (^(fi15))
                }
            }
            _ = v16
        } else {
            _ = len(v4)
        }
        _ = fi15
    }
    v13 += v13
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
