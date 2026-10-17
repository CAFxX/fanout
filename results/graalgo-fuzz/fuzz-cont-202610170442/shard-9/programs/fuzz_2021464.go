package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0(p0 int32) int8 {
    v6 := GenS1{F0: 0, F1: p0, F2: (true && (!(!((true && true) && false))))}
    if false {
        var v7 uint32 = 4294967295
        _ = v7
    }
    v8 := []int{}
    _ = p0
    _ = v6
    _ = v8
    return 86
}
func genFn1() int16 {
    var v9 int32 = -2147483648
    var v10 uint64 = 0
    v11 := GenS2{F0: (0.1 > 1e10), F1: (v9 + (v9 & (v9 % ((v9) | 1))))}
    _ = v10
    _ = v11
    _ = v9
    return 935
}

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 4294967295
    v3 := make(map[string]string)
    v4 := make(map[uint32]bool)
    var v5 uint64 = 2
    v2 = (v2 % ((v2) | 1))
    var v12 int64 = -9223372036854775808
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS3{F0: 74, F1: 42}
        _ = fi13
        _ = v14
    }
    fmt.Println(v1, v1)
    v1 = v1
    if (true && false) {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int16 = 1000
            for fi17 := 0; fi17 < 13; fi17++ {
                v5 = (v5 / (((^(v5))) | 1))
                var v18 uint16 = 906
                delete(v4, (v2 << ((0) & 63)))
                _ = fi17
                _ = v18
            }
            v12 = (v12 / ((v12) | 1))
            _ = fi15
            _ = v16
        }
        v5 -= 778
        fmt.Println(v5, v1)
    }
    v19 := GenS3{F0: -857, F1: 32767}
    v20 := []int{0, 32768}
    v1 += float32(0.1)
    fmt.Println(v1, v5, v12)
    v21 := GenS1{F0: (^(v5)), F1: 1, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
