package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 int16) int64 {
    p0 = 42
    if true {
        var v6 int32 = -1
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := []int{32767, fi7}
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    if true {
        switch "" {
        case "a":
            var v9 int32 = 0
            _ = v9
        default:
            if true {
                p0 = p0
                v10 := GenS1{F0: p0}
                var v11 int64 = 1
                _ = v10
                _ = v11
            }
        }
    }
    var v12 uint64 = 4294967295
    _ = p0
    _ = v12
    return 0
}

func main() {
    var v1 int8 = -1
    var v2 uint16 = 65535
    v3 := make(map[int]uint32)
    v4 := []uint32{612, 2147483647}
    v5 := make(map[uint32]string)
    var v13 float32 = 1e10
    v13 += v13
    v14 := GenS1{F0: 1000}
    fmt.Println(v2, v1, v2)
    v2 = 2
    v15 := GenS2{F0: "hello", F1: v2}
    var v16 float32 = 1.5
    fmt.Println(v2, v1)
    if (!(32767 > 0)) {
        switch -1 {
        case 111:
            fmt.Println(v1, v16)
        case -9223372036854775808:
            v2 = v2
        default:
            v17 := GenS1{F0: 1000}
            _ = v17
        }
        v3[42] = 1
    } else {
        fmt.Println(v2, v2, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
