package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }
func genFn0(p0 int64) uint32 {
    p0 += p0
    var v6 float32 = 1.5
    if (!(0.0 != v6)) {
        var v7 int8 = 79
        var v8 int = 0
        _ = v7
        _ = v8
    }
    v6 -= 0.0
    _ = p0
    _ = v6
    return 2
}

func main() {
    var v1 float64 = float64(1.5)
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 uint32 = 2147483648
    v5 := []int{336, 0, -148}
    if (100 != -12) {
        v1 = v1
    } else {
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(fi9, v4, v1)
            _ = fi9
        }
        fmt.Println(v4)
    }
    fmt.Println(v1)
    fmt.Println(v4, v1)
    if false {
        v10 := GenS1{F0: 100}
        v1 = float64(0.1)
        _ = v10
    } else {
        if (774 == 119) {
            v11 := v3["hello"]
            _ = v11
        }
        var v12 uint64 = 0
        _ = v12
    }
    v13 := GenS2{F0: 949, F1: 49, F2: 818}
    var v14 int16 = 32767
    for fi15 := 0; fi15 < 12; fi15++ {
        if false {
            v16 := GenS3{F0: -719, F1: 9223372036854775807, F2: 928}
            if (!false) {
                fmt.Println(v4, v14)
            } else {
                v2[(fi15 * (fi15 ^ 32768))] = (fi15 | 618)
                var v17 float32 = (3.14159 + -2.25)
                _ = v17
            }
            _ = v16
        } else {
            delete(v2, (-(fi15)))
            var v18 uint32 = (v4 + v4)
            _ = v18
        }
        if len(v5) > 0 {
            v5[0] = fi15
        }
        _ = fi15
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
