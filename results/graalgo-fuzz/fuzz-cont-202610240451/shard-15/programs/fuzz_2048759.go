package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint8, p1 uint, p2 int8) int32 {
    var v4 uint32 = 2147483648
    for fi5 := 0; fi5 < 12; fi5++ {
        v4 = 1
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: p0, F1: -67, F2: 11}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 uint64 = 42
    for fi9 := 0; fi9 < 14; fi9++ {
        p0 = (^(p0))
        if true {
            if (2 <= 42) {
                for fi10 := 0; fi10 < 15; fi10++ {
                    var v11 int8 = p2
                    _ = fi10
                    _ = v11
                }
                v12 := []int{-9223372036854775808, (fi9 | fi9)}
                _ = v12
            }
            v13 := []int{(-(fi9)), fi9, 906}
            _ = v13
        } else {
            var v14 float64 = (-2.25 - 1e10)
            _ = v14
        }
        v15 := []int{fi9}
        _ = fi9
        _ = v15
    }
    v16 := GenS2{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v4
    _ = v8
    return 2147483647
}
func genFn1(p0 int32, p1 uint8, p2 int64) int64 {
    v17 := GenS2{F0: 42}
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return p2
}

func main() {
    var v1 uint32 = 641
    var v2 int = -1
    var v3 int16 = 0
    v18 := GenS2{F0: -128}
    if (v1 >= 1) {
        var v19 int64 = -1
        if false {
            v20 := []int{v2, (v2 << ((42) & 63)), 32768, v2}
            fmt.Println(v2, v19, v2)
            _ = v20
        }
        _ = v19
    } else {
        if (!false) {
            v3 = (v3 + v3)
        } else {
            v3 = v3
            v21 := GenS2{F0: -128}
            _ = v21
        }
        v22 := GenS1{F0: 255, F1: 0, F2: 42}
        _ = v22
    }
    v3 += v3
    switch (v2 % (((v2 << ((0) & 63))) | 1)) {
    case 0:
        v23 := GenS1{F0: 255, F1: -2147483648, F2: 778}
        _ = v23
    default:
        if false {
            v24 := GenS2{F0: -77}
            var v25 uint64 = 18446744073709551615
            _ = v24
            _ = v25
        }
    }
    var v26 int16 = (v3 - -32768)
    var v27 uint64 = 288
    if true {
        var v28 float32 = 0.1
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint16 = 852
            for fi31 := 0; fi31 < 11; fi31++ {
                if (!true) {
                    fmt.Println(fi29, v30, v2)
                    v1 = (v1 * uint32(v30))
                }
                _ = fi31
            }
            v3 -= v3
            _ = fi29
            _ = v30
        }
        v1 += 2147483647
        _ = v28
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 uint64 = 0
        _ = fi32
        _ = v33
    }
    v34 := []int{(v2 | v2), (v2 & 42), (^(v2))}
    var v35 float64 = -2.25
    switch -991 {
    case -1:
        var v36 int64 = 1
        _ = v36
    case -817:
        if len(v34) > 0 {
            v34[0] = v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
}
