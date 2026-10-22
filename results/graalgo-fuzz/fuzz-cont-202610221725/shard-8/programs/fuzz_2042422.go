package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint, p1 uint32, p2 uint) uint32 {
    switch 32767 {
    case -687:
        v7 := GenS1{F0: -1, F1: -1, F2: false}
        _ = v7
    default:
        if false {
            p0 = p2
            p1 = (p1 * 65535)
            v8 := GenS2{F0: 0}
            _ = v8
        } else {
            v9 := []int{}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        switch "\n\t" {
        case "hello":
            v11 := []int{}
            _ = v11
        }
        v12 := GenS2{F0: -1}
        _ = fi10
        _ = v12
    }
    v13 := []int{2147483647, -2147483648, 9223372036854775807, 0}
    var v14 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return (p1 * (p1 + 2))
}

func main() {
    var v1 uint64 = 44
    var v2 int = 4294967295
    var v3 uint32 = 0
    var v4 int = (v2 / ((v2) | 1))
    v5 := make(map[uint32]string)
    var v6 int16 = -666
    v6 = 0
    for fi15 := 0; fi15 < 7; fi15++ {
        if false {
            v16 := []int{32767, v4}
            for fi17 := 0; fi17 < 8; fi17++ {
                if false {
                    v16 = append(v16, -169)
                    if (true || false) {
                        var v18 uint64 = (v1 & 925)
                        var v19 uint8 = 127
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint16 = 889
                        _ = v20
                    }
                    v21 := GenS2{F0: 0}
                    fmt.Println(v4, fi15, v3)
                    v3 = v3
                    _ = v21
                }
                v1 -= (v1 ^ v1)
                if (!(255 <= 1)) {
                    v1 = 18446744073709551615
                }
                _ = fi17
            }
            _ = v16
        } else {
            v22 := GenS2{F0: -39}
            v2 = (v4 & -867)
            var v23 uint = 2
            _ = v22
            _ = v23
        }
        _ = fi15
    }
    v1 = 18446744073709551615
    fmt.Println(v2, v2, v6)
    var v24 uint32 = (^(v3))
    v25 := GenS1{F0: -1, F1: 1, F2: false}
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := GenS2{F0: 896}
        var v28 int64 = 1
        _ = fi26
        _ = v27
        _ = v28
    }
    v5[4294967295] = "foo bar"
    v6 = v6
    v29 := []int{(v2 - v4)}
    var v30 int16 = v6
    for fi31 := 0; fi31 < 11; fi31++ {
        v30 -= 42
        v6 -= v6
        _ = fi31
    }
    var v32 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
