package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
func genFn0() uint8 {
    var v4 int8 = 100
    var v5 bool = (!true)
    _ = v4
    _ = v5
    return 128
}
func genFn1(p0 uint32, p1 string, p2 bool) int {
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 = p0
        _ = fi6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = (p0 << ((42) & 63))
        _ = fi7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        switch 42 {
        case 9223372036854775807:
            var v9 uint16 = 0
            _ = v9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    v1 := []int{0, 32767}
    var v2 uint64 = 42
    var v3 int8 = int8(v2)
    fmt.Println(v2, v3)
    if false {
        fmt.Println(v3)
        v1 = append(v1, 203)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 uint8 = 246
    v11 := GenS1{F0: 1}
    if ((v2 & 2) > v2) {
        for fi12 := 0; fi12 < 17; fi12++ {
            v10 = (v10 ^ (^(v10)))
            fmt.Println(fi12, v3)
            fmt.Println(v3, v10, v2)
            _ = fi12
        }
        if true {
            v2 = v2
        }
        v1 = append(v1, 743)
    } else {
        v2 = v2
        if false {
            var v13 uint64 = 18446744073709551615
            if false {
                switch -2147483648 {
                case 0:
                    v10 -= v10
                }
            } else {
                v14 := GenS2{F0: v3}
                _ = v14
            }
            _ = v13
        }
    }
    fmt.Println(v10, v10)
    v15 := GenS1{F0: v3}
    if len(v1) > 0 {
        v1[0] = 994
    }
    v16 := GenS1{F0: v3}
    v17 := GenS2{F0: -128}
    v10 = uint8(v2)
    var v18 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
