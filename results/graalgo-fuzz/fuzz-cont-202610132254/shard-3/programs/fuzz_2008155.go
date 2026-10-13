package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 int32, p1 int64) int {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{(fi4 - 2147483647), fi4, (-(fi4)), (fi4 | fi4)}
        v6 := GenS1{F0: 1, F1: 2147483648, F2: 303}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int16 = 42
    p0 = 891
    _ = p0
    _ = p1
    _ = v7
    return 32767
}
func genFn1(p0 uint32) uint32 {
    var v8 uint8 = 0
    var v9 uint64 = 4294967295
    v10 := GenS2{F0: 445, F1: 1}
    var v11 uint64 = (^(v9))
    var v12 int8 = -1
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{2147483647, 910}
    v2 := make(map[int]string)
    var v3 uint32 = 42
    var v13 bool = (32767 == 0)
    var v14 uint64 = 553
    delete(v2, 1)
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(v14)
        v16 := v2[(fi15 | fi15)]
        _ = fi15
        _ = v16
    }
    switch 2147483647 {
    case 127:
        v3 = (v3 << ((255) & 63))
    case -9223372036854775808:
        var v17 uint32 = 2147483647
        _ = v17
    }
    v13 = (0 <= 65535)
    switch 876 {
    case -1:
        v14 -= 161
    case 541:
        var v18 float64 = -2.25
        _ = v18
    case 32768:
        fmt.Println(v13, v14, v14)
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        v13 = v13
        if (206 < 42) {
            var v20 int32 = 1
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := GenS1{F0: 2147483648, F1: v3, F2: 979}
                _ = fi21
                _ = v22
            }
            v23 := GenS1{F0: v3, F1: (^(v3)), F2: 2}
            _ = v20
            _ = v23
        }
        _ = fi19
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 int64 = -612
        _ = fi24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
