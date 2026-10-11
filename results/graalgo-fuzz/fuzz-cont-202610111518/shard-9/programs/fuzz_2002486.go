package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 string }
func genFn0(p0 uint64, p1 int64) int32 {
    v7 := GenS1{F0: 0, F1: 913}
    if true {
        var v8 int32 = 2147483647
        p0 = (p0 | p0)
        _ = v8
    } else {
        p0 = p0
        v9 := []int{1}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 int64) uint32 {
    if false {
        var v10 uint32 = 65535
        if true {
            v11 := []int{-693}
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 float32 = ((1e10 - 0.1) - 1e10)
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            fmt.Println(v10)
            var v14 uint64 = 42
            for fi15 := 0; fi15 < 16; fi15++ {
                p1 += (p1 >> ((42) & 63))
                v16 := []int{(fi15 / ((fi15) | 1))}
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        v17 := []int{2147483647}
        _ = v10
        _ = v17
    } else {
        if true {
            p1 = p1
            if true {
                v18 := GenS2{F0: (p1 != (p1 / (((p1 * p1)) | 1))), F1: 42}
                _ = v18
            }
            var v19 int32 = 0
            _ = v19
        }
        var v20 int64 = 0
        var v21 bool = (false && (1 < 0))
        _ = v20
        _ = v21
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 688
}

func main() {
    var v1 float32 = 1.5
    var v2 uint16 = 664
    var v3 uint32 = 168
    v4 := make(map[int]int)
    var v5 uint16 = (v2 | v2)
    var v6 int8 = -128
    delete(v4, 1)
    var v22 uint16 = v2
    v1 = 0.0
    v23 := GenS2{F0: true, F1: 0}
    if (!true) {
        v1 = (((v1 + v1) - (0.1 - v1)) * v1)
    } else {
        for fi24 := 0; fi24 < 14; fi24++ {
            v25 := GenS1{F0: 42, F1: 570}
            _ = fi24
            _ = v25
        }
        v6 += v6
    }
    v5 = 716
    v22 += 0
    v26 := GenS3{F0: (("" + "\n\t") + "hello")}
    v2 -= 1
    var v27 uint64 = 4294967295
    var v28 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
