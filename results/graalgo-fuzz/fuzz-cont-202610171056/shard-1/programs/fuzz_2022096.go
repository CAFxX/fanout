package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 string, p1 uint32, p2 uint64) uint8 {
    v7 := []int{17, 14, 718}
    if (true && true) {
        fmt.Println(p0, p1, p2)
        fmt.Println(p0)
        if (true && (!(0 < 18446744073709551615))) {
            var v8 uint32 = 810
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint32 = p1
        _ = fi9
        _ = v10
    }
    _ = len(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 7
}
func genFn1(p0 uint64, p1 int16) int16 {
    var v11 float64 = 3.14159
    var v12 float64 = -2.25
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return p1
}

func main() {
    var v1 bool = false
    var v2 uint16 = 522
    var v3 uint16 = v2
    var v4 int8 = 127
    v5 := make(map[uint32]string)
    var v6 int32 = 370
    if true {
        fmt.Println(v4)
    }
    var v13 uint8 = 0
    v14 := []int{9223372036854775807, int(v2)}
    for fi15 := 0; fi15 < 9; fi15++ {
        if (-1 < 1) {
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 int64 = -586
                v18 := GenS1{F0: 65535}
                _ = fi16
                _ = v17
                _ = v18
            }
            v3 -= (v2 << ((2) & 63))
            var v19 int16 = 1
            _ = v19
        }
        v1 = (2147483648 != 42)
        var v20 int32 = -709
        _ = fi15
        _ = v20
    }
    if false {
        v6 = v6
        v21 := GenS1{F0: 0}
        _ = v21
    }
    v14 = append(v14, -30)
    fmt.Println(v4, v3)
    v14 = append(v14, int(v13))
    v22 := GenS2{F0: -958}
    var v23 uint16 = 506
    v24 := GenS2{F0: -1}
    if true {
        var v25 int = -803
        fmt.Println(v3, v13)
        var v26 int8 = (v4 % ((v4) | 1))
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
