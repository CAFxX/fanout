package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }

func main() {
    var v1 int8 = -1
    var v2 uint64 = 969
    var v3 float64 = (3.14159 * float64(0.1))
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 uint8 = 0
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 uint32 = 42
            _ = fi6
            _ = v7
        }
        switch 9223372036854775807 {
        case -232:
            var v8 uint32 = 2
            _ = v8
        case -9223372036854775808:
            v9 := GenS1{F0: 2, F1: -1}
            _ = v9
        }
        _ = fi4
        _ = v5
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v10 uint64 = v2
        _ = v10
    case 0:
        if true {
            var v11 float32 = (((3.14159 * 1e10) - float32(v3)) + 0.1)
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 bool = true
        v14 := []int{(fi12 >> ((0) & 63)), fi12, 764}
        _ = fi12
        _ = v13
        _ = v14
    }
    switch -560 {
    case -1:
        v15 := []int{32768, -2147483648}
        _ = v15
    default:
        v1 -= (v1 >> ((829) & 63))
    }
    fmt.Println(v1, v3)
    var v16 uint32 = 4294967295
    if (v16 != uint32(v2)) {
        var v17 bool = false
        _ = v17
    } else {
        var v18 int32 = -1
        var v19 float32 = ((-2.25 + 0.0) * 3.14159)
        _ = v18
        _ = v19
    }
    if (false && (false || true)) {
        v2 -= 0
        v16 = v16
    } else {
        var v20 float64 = (v3 + (0.0 * (v3 * 0.1)))
        v21 := GenS1{F0: 0, F1: 0}
        for fi22 := 0; fi22 < 14; fi22++ {
            if ((!(-301 <= 42)) && true) {
                v23 := GenS1{F0: 255, F1: int64(v2)}
                var v24 uint = 895
                _ = v23
                _ = v24
            } else {
                if ((-1 < 42) && ((((3.14159 / v3) / 3.14159) - 0.0) > v3)) {
                    var v25 int32 = -1
                    v16 += v16
                    v25 = 2147483647
                    _ = v25
                }
                v26 := []int{fi22, fi22, 128}
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = v26
            }
            _ = fi22
        }
        _ = v20
        _ = v21
    }
    v27 := []int{127, -9223372036854775808, -724}
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 uint32 = (^(v16))
        v29 += 2147483647
        _ = fi28
        _ = v29
    }
    v3 = ((v3 - v3) / (v3 + v3))
    v2 -= (v2 * v2)
    for fi30 := 0; fi30 < 1; fi30++ {
        v2 -= 1
        if ((v2 & v2) < 0) {
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
        } else {
            if len(v27) > 0 {
                v27[0] = (-(fi30))
            }
            v3 += v3
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v3
}
