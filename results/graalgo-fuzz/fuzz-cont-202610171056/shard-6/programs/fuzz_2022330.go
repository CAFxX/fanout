package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int32; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 uint32) string {
    var v4 int64 = 42
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: -39}
        if false {
            if false {
                for fi7 := 0; fi7 < 5; fi7++ {
                    v4 -= 42
                    _ = fi7
                }
            } else {
                for fi8 := 0; fi8 < 8; fi8++ {
                    if (false || true) {
                        var v9 uint32 = p0
                        _ = v9
                    }
                    _ = fi8
                }
            }
        }
        var v10 float32 = (-2.25 * 0.1)
        _ = fi5
        _ = v10
        _ = v6
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        if false {
            v12 := GenS1{F0: 92}
            fmt.Println(v4, p0, fi11)
            _ = v12
        }
        p0 -= (p0 | (p0 * p0))
        for fi13 := 0; fi13 < 6; fi13++ {
            switch fi13 {
            case -591:
                v14 := []int{fi11, 4294967295}
                _ = v14
            case 1:
                var v15 int16 = -32768
                _ = v15
            case -2147483648:
                v16 := []int{(fi11 - fi11), -2147483648, (fi13 / (((fi13 / ((fi11) | 1))) | 1)), fi11}
                _ = v16
            default:
                var v17 float32 = 1e10
                _ = v17
            }
            _ = fi13
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        if false {
            p0 = p0
        } else {
            v19 := GenS1{F0: -1}
            _ = v19
        }
        _ = fi18
    }
    _ = p0
    _ = v4
    return ""
}
func genFn1(p0 uint64) uint16 {
    var v20 float64 = 3.14159
    v21 := GenS2{F0: 175}
    p0 = 2
    var v22 uint = 2
    v20 = (1e10 + v20)
    _ = p0
    _ = v20
    _ = v21
    _ = v22
    return 32767
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 42
    v3 := []uint32{}
    v23 := GenS1{F0: 100}
    _ = len(v3)
    v2 += v2
    var v24 int = -9223372036854775808
    v25 := GenS3{F0: ((v24 << ((42) & 63)) < v24), F1: -468, F2: 0}
    fmt.Println(v2, v24)
    for fi26 := 0; fi26 < 4; fi26++ {
        v2 -= v2
        if len(v3) > 0 {
            v3[0] = v2
        }
        switch (v24 & (v24 / ((v24) | 1))) {
        case -1:
            fmt.Println(fi26)
        case 42:
            var v27 int16 = 336
            _ = v27
        }
        _ = fi26
    }
    v28 := GenS4{F0: (v2 << ((18446744073709551615) & 63)), F1: 4294967295}
    var v29 bool = false
    if len(v3) > 0 {
        v3[0] = (v2 - v2)
    }
    v29 = (!(!(42 > 0)))
    var v30 int16 = -244
    var v31 bool = (!v29)
    v24 = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
