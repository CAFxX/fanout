package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int; F2 int16 }
func genFn0() int32 {
    var v6 float64 = float64(3.14159)
    var v7 uint32 = 427
    _ = v6
    _ = v7
    return -36
}
func genFn1(p0 int64, p1 uint8) string {
    v8 := []int{}
    p0 = p0
    _ = p0
    _ = p1
    _ = v8
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 255
    v3 := []int{-1}
    var v4 float64 = (0.1 * ((-2.25 / (((-2.25 + -2.25) / -2.25) + (0.0 - 1e10))) * 1e10))
    var v5 uint64 = 18446744073709551615
    if true {
        var v9 float64 = v4
        if false {
            v10 := GenS1{F0: 4294967295, F1: 2147483647, F2: -128}
            v4 = v4
            _ = v10
        }
        _ = v9
    } else {
        if (128 == 32767) {
            if len(v3) > 0 {
                v3[0] = -1
            }
            v11 := v1[-2147483648]
            v12 := v1[-414]
            _ = v11
            _ = v12
        } else {
            var v13 string = (("hello" + "hello") + (("" + "hello") + ("" + "a")))
            v5 += 4294967295
            var v14 int = 9223372036854775807
            _ = v13
            _ = v14
        }
        v5 += v5
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4)
    for fi15 := 0; fi15 < 20; fi15++ {
        v4 = float64((0.0 + 0.0))
        fmt.Println(v5, v5)
        _ = fi15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 int8 = -99
        var v18 uint64 = (v5 ^ (v5 + v5))
        var v19 int = (fi16 & (fi16 ^ 2147483647))
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 uint64 = 42
    v2 -= (v2 | v2)
    _ = len(v3)
    v2 = 32767
    if len(v3) > 0 {
        v3[0] = 908
    }
    if false {
        v20 = (v5 | (v5 | v20))
    } else {
        for fi21 := 0; fi21 < 10; fi21++ {
            fmt.Println(v2, fi21)
            if true {
                var v22 uint64 = v20
                _ = v22
            }
            v3 = append(v3, -9223372036854775808)
            _ = fi21
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            if (!false) {
                var v24 float64 = 1.5
                var v25 float32 = (3.14159 / (float32(v24) + float32(((v24 * v4) * (v4 + v24)))))
                _ = v24
                _ = v25
            } else {
                v26 := GenS1{F0: 389, F1: fi23, F2: 100}
                v2 -= (v2 ^ (v2 & 255))
                _ = v26
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                v5 = v5
                fmt.Println(v2, v5)
                _ = fi27
            }
            _ = fi23
        }
        v28 := GenS2{F0: 857}
        _ = v28
    }
    v5 += v20
    for fi29 := 0; fi29 < 10; fi29++ {
        v5 -= v20
        var v30 int64 = 1
        v31 := GenS3{F0: (!(((v30 / (((v30 | v30)) | 1)) <= (v30 << ((0) & 63))) || (1 <= 152))), F1: (fi29 | 32767), F2: 42}
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
