package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -732
    v3 := make(map[string]bool)
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: 255}
        delete(v1, 91)
        v6 := GenS2{F0: -939, F1: 100}
        _ = fi4
        _ = v5
        _ = v6
    }
    if true {
        v7 := GenS1{F0: 4294967295}
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        _ = len(v1)
        var v9 int32 = -717
        if true {
            v10 := GenS2{F0: 77, F1: 1}
            fmt.Println(fi8)
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    v2 = (v2 % ((-2147483648) | 1))
    var v11 float32 = 0.0
    var v12 uint64 = 814
    if ((2 < 2147483647) || false) {
        var v13 int16 = 0
        var v14 float32 = (v11 + (0.0 + v11))
        _ = v13
        _ = v14
    }
    fmt.Println(v12)
    for fi15 := 0; fi15 < 8; fi15++ {
        if ((fi15 / (((^(v2))) | 1)) >= (v2 / (((v2 - -522)) | 1))) {
            v16 := GenS1{F0: 0}
            _ = v16
        } else {
            var v17 bool = ((1 > 42) && true)
            for fi18 := 0; fi18 < 2; fi18++ {
                v12 -= v12
                _ = fi18
            }
            v12 += 4294967295
            _ = v17
        }
        if false {
            v11 += 0.1
        }
        _ = fi15
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        v2 = (v2 & v2)
        v2 += (fi19 ^ (fi19 >> ((200) & 63)))
        var v20 uint32 = 387
        _ = fi19
        _ = v20
    }
    v21 := []int{v2, (^(v2)), 621}
    v22 := GenS3{F0: "hello"}
    var v23 int64 = -9223372036854775808
    v1[354] = (v2 ^ -98)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
