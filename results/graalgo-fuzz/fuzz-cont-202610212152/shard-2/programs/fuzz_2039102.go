package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int }
func genFn0(p0 int8, p1 uint32, p2 int64) uint32 {
    p1 = (^(p1))
    p1 = p1
    var v7 int = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 % (((^(p1))) | 1))
}
func genFn1(p0 uint32, p1 uint16) uint64 {
    var v8 int = 906
    v9 := GenS1{F0: 18446744073709551615, F1: 32767}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int8 = 42
    var v3 bool = true
    var v4 int8 = (v2 | v2)
    var v5 uint8 = 128
    v6 := make(map[uint32]bool)
    var v10 uint8 = (v5 / ((v5) | 1))
    var v11 int8 = v4
    v5 += (v5 - v10)
    var v12 uint = 0
    fmt.Println(v3, v11, v10)
    fmt.Println(v11, v1)
    for fi13 := 0; fi13 < 5; fi13++ {
        if false {
            v14 := GenS2{F0: 116, F1: 65535, F2: (fi13 >> ((v12) & 63))}
            v15 := GenS1{F0: genFn1(uint32(v12), 255), F1: 659}
            _ = v14
            _ = v15
        } else {
            var v16 int8 = v4
            _ = v16
        }
        _ = fi13
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 int = (fi17 & 269)
        var v19 int32 = 42
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := []int{v1}
    if true {
        if (!true) {
            if v3 {
                v21 := GenS2{F0: 32767, F1: 2, F2: (v1 << ((2) & 63))}
                _ = v21
            } else {
                v3 = v3
            }
        }
    } else {
        v1 = (v1 | 0)
    }
    if v3 {
        _ = len(v6)
    } else {
        var v22 uint32 = 0
        _ = v22
    }
    if v3 {
        fmt.Println(v3, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
