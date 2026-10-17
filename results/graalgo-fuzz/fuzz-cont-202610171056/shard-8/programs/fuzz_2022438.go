package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 int; F2 uint }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0(p0 int64, p1 uint16) uint32 {
    v6 := []int{499, -1000, 127}
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 4294967295, F1: 54, F2: p0}
        fmt.Println(p0, fi7, p1)
        _ = fi7
        _ = v8
    }
    p0 = -9223372036854775808
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint8 = 2
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{2147483648, 1, 4294967295}
    var v3 int16 = -1
    var v4 uint32 = 4294967295
    v5 := make(map[string]int)
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(fi11, fi11, v3)
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS1{F0: 18446744073709551615, F1: 63, F2: -9223372036854775808}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v14 := GenS1{F0: 18446744073709551615, F1: 161, F2: -797}
    var v15 uint = 375
    v16 := GenS2{F0: -2147483648, F1: -318, F2: (v15 / ((1) | 1))}
    var v17 bool = false
    var v18 int8 = 1
    for fi19 := 0; fi19 < 4; fi19++ {
        v15 -= v15
        _ = fi19
    }
    v20 := v1["\n\t"]
    fmt.Println(v20, v18, v15)
    v21 := GenS3{F0: -1, F1: v20}
    fmt.Println(v4, v17, v17)
    v15 = v15
    if ((!v17) || ((-2.25 * 1e10) <= float32(1e10))) {
        v4 = genFn0(int64(v15), 2)
        for fi22 := 0; fi22 < 13; fi22++ {
            fmt.Println(v4, fi22, v15)
            _ = fi22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
