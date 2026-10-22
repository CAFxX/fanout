package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
func genFn0() int8 {
    var v7 float64 = (0.0 * (0.0 + float64((1.5 / -2.25))))
    fmt.Println(v7)
    var v8 int = 128
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 string, p1 int32) int16 {
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            if ((-(p1)) <= p1) {
                var v11 uint64 = 160
                p0 = (p0 + ("foo bar" + p0))
                _ = v11
            }
            fmt.Println(fi10, fi10)
            _ = fi10
        }
        _ = fi9
    }
    if true {
        var v12 int8 = 100
        var v13 uint = 447
        var v14 float32 = 1e10
        _ = v12
        _ = v13
        _ = v14
    }
    fmt.Println(p1)
    p1 += (-(p1))
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return -32768
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{32767, 0, 1, 119}
    var v3 uint16 = 398
    var v4 int8 = 42
    var v5 uint8 = 42
    var v6 int8 = v4
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS1{F0: 65535}
        v3 = (v3 % (((v3 * 32767)) | 1))
        var v17 uint16 = v3
        _ = fi15
        _ = v16
        _ = v17
    }
    v5 -= (v5 << ((4294967295) & 63))
    v18 := GenS1{F0: 1}
    fmt.Println(v4, v6)
    var v19 float64 = (float64((0.1 - -2.25)) * ((3.14159 * 0.1) + (-2.25 - (0.0 - ((1e10 * float64(0.1)) * 1e10)))))
    for fi20 := 0; fi20 < 3; fi20++ {
        if (((v3 / ((2) | 1)) < (v3 << ((0) & 63))) || (true && true)) {
            var v21 int8 = genFn0()
            var v22 int64 = -9223372036854775808
            _ = v21
            _ = v22
        }
        var v23 uint8 = (v5 | (v5 / ((v5) | 1)))
        v6 -= (v6 + v6)
        _ = fi20
        _ = v23
    }
    if (((1 >= (^(v3))) || true) || (false || false)) {
        v3 = (v3 + v3)
        v5 += (v5 * v5)
    }
    var v24 uint16 = (v3 - (^(v3)))
    var v25 bool = false
    v26 := GenS2{F0: 439}
    v27 := v1[4294967295]
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
