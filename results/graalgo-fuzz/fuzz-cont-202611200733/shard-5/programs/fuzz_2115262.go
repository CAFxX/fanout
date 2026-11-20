package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int }
type GenS2 struct { F0 uint16; F1 uint; F2 uint64 }
func genFn0(p0 uint32) int8 {
    var v6 int = 128
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint = 65535
        var v9 uint32 = 42
        v10 := []int{v6}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if true {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int32 = 1
            _ = fi11
            _ = v12
        }
        var v13 float64 = (float64((1e10 / 1.0)) - (float64(3.14159) / 1.5))
        fmt.Println(v13, v6)
        _ = v13
    }
    _ = p0
    _ = v6
    return -1
}

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    var v3 int8 = -128
    var v4 bool = (!v1)
    v5 := make(map[int]bool)
    var v14 float32 = float32(((1.5 + 0.1) * -2.25))
    v3 += (v3 * 127)
    v1 = (!false)
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    v17 := []int{0, 442}
    v18 := GenS1{F0: 2147483648, F1: 1, F2: 888}
    var v19 int = -9223372036854775808
    for fi20 := 0; fi20 < 19; fi20++ {
        v19 += -188
        var v21 int = (v19 - v19)
        _ = len(v5)
        _ = fi20
        _ = v21
    }
    if (!v1) {
        v22 := GenS1{F0: 0, F1: 224, F2: 32767}
        v19 -= 127
        _ = v22
    }
    if v4 {
        if len(v17) > 0 {
            v17[0] = v19
        }
        var v23 int16 = 42
        _ = v23
    }
    var v24 int64 = -692
    for fi25 := 0; fi25 < 20; fi25++ {
        for fi26 := 0; fi26 < 18; fi26++ {
            v24 = (v24 + (v24 & -877))
            if (!(v19 < (v19 % ((-324) | 1)))) {
                fmt.Println(v1, v14)
                var v27 uint64 = 2
                _ = v27
            } else {
                v17 = append(v17, fi26)
                var v28 uint32 = 0
                delete(v5, fi26)
                v4 = (1 < 1)
                var v29 int64 = (^(v24))
                _ = v28
                _ = v29
            }
            _ = fi26
        }
        for fi30 := 0; fi30 < 7; fi30++ {
            v31 := GenS2{F0: 100, F1: 42, F2: 44}
            var v32 uint32 = 0
            _ = fi30
            _ = v31
            _ = v32
        }
        v19 -= fi25
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
