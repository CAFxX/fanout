package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int) uint16 {
    fmt.Println(p0, p0)
    var v4 uint64 = 1
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 18446744073709551615, F1: uint16(p0), F2: 42}
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 621
}
func genFn1() uint8 {
    switch "" {
    case "\n\t":
        var v8 uint8 = 127
        _ = v8
    default:
        v9 := GenS2{F0: 2}
        _ = v9
    }
    var v10 int = -856
    v11 := []int{(v10 * (v10 ^ (v10 * v10))), v10}
    v11 = append(v11, v10)
    v11 = append(v11, 127)
    _ = v10
    _ = v11
    return 255
}

func main() {
    v1 := []string{}
    var v2 uint32 = 891
    var v3 int64 = -9223372036854775808
    var v12 int64 = (v3 - -9223372036854775808)
    v13 := GenS2{F0: 42}
    if true {
        v3 = (v3 % ((int64(v2)) | 1))
    }
    var v14 uint8 = 19
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint64 = 2
        v12 = v12
        _ = fi15
        _ = v16
    }
    v14 = 127
    v14 -= (v14 * v14)
    if (42 != 552) {
        v12 = 583
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
