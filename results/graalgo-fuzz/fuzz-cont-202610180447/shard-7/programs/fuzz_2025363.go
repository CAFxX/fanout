package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint32) uint {
    p0 -= 0
    p0 -= (p0 - (p0 ^ p0))
    var v5 int32 = -1
    fmt.Println(v5)
    if true {
        if true {
            v5 -= v5
        }
        v5 = (v5 << ((42) & 63))
        fmt.Println(p0)
    }
    _ = p0
    _ = v5
    return 853
}
func genFn1(p0 bool, p1 uint32, p2 uint) int8 {
    v6 := []int{-337}
    var v7 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int16 = -1
    var v2 int64 = 0
    v3 := []int{9223372036854775807, -2147483648, -9223372036854775808, -2147483648}
    var v4 bool = (1 != -2147483648)
    v8 := GenS1{F0: 0, F1: 0, F2: (v2 - (v2 << ((genFn0(2147483647)) & 63)))}
    v9 := GenS2{F0: (75 == 822), F1: 0}
    var v10 int8 = genFn1(false, 42, 65535)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v11 float64 = 0.0
    v1 += 42
    if v4 {
        var v12 uint64 = 42
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v12
    } else {
        v1 = (v1 & -716)
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint32 = 2147483647
        _ = fi13
        _ = v14
    }
    if (v4 || (!(v11 == v11))) {
        for fi15 := 0; fi15 < 13; fi15++ {
            switch (fi15 ^ (fi15 / ((fi15) | 1))) {
            case 9223372036854775807:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi15
        }
        _ = len(v3)
        fmt.Println(v4)
    } else {
        v4 = true
        if v4 {
            switch ("a" + ("hello" + "\n\t")) {
            case "\n\t":
                var v16 int = 0
                _ = v16
            case "":
                v4 = true
            }
            v17 := GenS2{F0: (-2147483648 < 42), F1: 127}
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
