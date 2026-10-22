package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0(p0 bool) uint16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint8 = 42
        _ = fi7
        _ = v8
    }
    p0 = true
    p0 = (9223372036854775807 != -549)
    _ = p0
    return 884
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]uint32)
    var v5 uint32 = 307
    var v6 int32 = -1
    if (!true) {
        v6 += (v6 & (v6 + 2147483647))
        fmt.Println(v6)
        v3 -= v3
    }
    if true {
        fmt.Println(v3, v5)
    }
    fmt.Println(v3, v3, v6)
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v5 -= 2
            v5 += v5
            v3 = (v3 * (v3 >> ((0) & 63)))
            _ = len(v4)
            fmt.Println(fi10, fi10)
            if true {
                v11 := []int{(^(fi10)), 4294967295, -93}
                fmt.Println(v5, v5, fi10)
                var v12 uint16 = 182
                _ = v11
                _ = v12
            } else {
                fmt.Println(v5)
                v5 += v5
                v3 = 1
                v5 = (v5 | (v5 & (v5 << ((4294967295) & 63))))
            }
            _ = fi10
        }
        v3 = -9223372036854775808
        _ = fi9
    }
    if false {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 int16 = -452
            switch fi13 {
            case 657:
                var v15 uint = 364
                _ = v15
            case 32768:
                v16 := GenS1{F0: 42, F1: -1}
                _ = v16
            case 4294967295:
                v14 = (v14 & int16(v6))
            default:
                fmt.Println(fi13, fi13, v6)
            }
            _ = fi13
            _ = v14
        }
        switch 0 {
        case -9223372036854775808:
            v17 := []int{int(v6), 330, 82}
            _ = v17
        }
    } else {
        fmt.Println(v6, v6, v3)
    }
    v6 = (v6 ^ v6)
    v3 += v3
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
