package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
func genFn0(p0 int64) int64 {
    v7 := GenS1{F0: true, F1: 0}
    p0 = 1
    var v8 uint16 = 32767
    p0 += p0
    v9 := []int{42, 128, -815, -2147483648}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 917
}
func genFn1(p0 uint64, p1 int32) uint16 {
    p0 = p0
    v10 := GenS1{F0: true, F1: 2}
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: true, F1: p0}
        p0 += p0
        v13 := []int{fi11, fi11, fi11, (fi11 | fi11)}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int8 = 0
    var v15 bool = true
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    _ = v15
    return 32767
}

func main() {
    var v1 int8 = -1
    var v2 uint32 = 596
    var v3 uint16 = 983
    v4 := make(map[string]bool)
    var v5 int64 = -1
    var v6 uint64 = 4294967295
    v16 := []int{-2147483648, -132}
    if (0.1 != 3.14159) {
        var v17 int = 127
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            switch (-(fi18)) {
            case 4294967295:
                v19 := GenS1{F0: (false || (108 < 255)), F1: 18446744073709551615}
                _ = v19
            case -117:
                v3 += v3
            default:
                fmt.Println(v5, v3)
            }
            _ = fi18
        }
    }
    v3 = 2
    if false {
        for fi20 := 0; fi20 < 15; fi20++ {
            if ((1 >= 127) || (genFn1(1, -1) <= v3)) {
                delete(v4, "")
                var v21 uint = 0
                for fi22 := 0; fi22 < 4; fi22++ {
                    if (!(!(0 == 32767))) {
                        v5 = (v5 ^ -1)
                    }
                    _ = fi22
                }
                _ = v21
            }
            v1 = v1
            _ = fi20
        }
        fmt.Println(v6, v1, v3)
        var v23 uint = 879
        _ = v23
    }
    v24 := GenS1{F0: true, F1: v6}
    v25 := GenS1{F0: (1 <= 4294967295), F1: (^(v6))}
    var v26 int32 = -304
    var v27 int64 = v5
    var v28 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
