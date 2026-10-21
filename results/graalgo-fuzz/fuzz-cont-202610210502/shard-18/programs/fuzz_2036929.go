package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 int64 }
func genFn0(p0 bool) int16 {
    v5 := GenS1{F0: 255, F1: 32767, F2: 65535}
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 65535, F1: 0, F2: 2}
            var v9 float32 = -2.25
            if true {
                v10 := []int{-315, (fi7 << ((4294967295) & 63))}
                var v11 uint64 = 28
                _ = v10
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        v12 := []int{(^(fi6)), -66, 128}
        v13 := GenS2{F0: 0, F1: 2147483648}
        _ = fi6
        _ = v12
        _ = v13
    }
    var v14 uint = 57
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS3{F0: "foo bar", F1: 255, F2: 0}
        if (-2147483648 != 164) {
            var v17 string = "a"
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = v14
    _ = v5
    return -936
}

func main() {
    v1 := []uint32{319}
    var v2 int8 = 1
    var v3 int = 0
    var v4 uint64 = 0
    v18 := GenS1{F0: 794, F1: 42, F2: 65535}
    fmt.Println(v2)
    v19 := GenS1{F0: 962, F1: 32767, F2: 885}
    var v20 float32 = 0.0
    if false {
        v21 := GenS2{F0: (v4 << ((42) & 63)), F1: 1}
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v2 -= v2
        v4 += v4
        v4 = 24
        _ = fi22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        _ = len(v1)
        _ = fi23
    }
    v3 += (^(v3))
    v3 = (v3 & v3)
    v3 += 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
