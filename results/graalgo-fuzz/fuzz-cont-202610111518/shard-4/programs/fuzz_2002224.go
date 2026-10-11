package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int32; F2 int8 }
func genFn0(p0 uint) uint16 {
    if ((true && false) && true) {
        p0 += (p0 & 0)
        p0 = 42
        var v6 int16 = 761
        _ = v6
    }
    p0 += (p0 | (p0 & (p0 % ((0) | 1))))
    fmt.Println(p0)
    var v7 int = 128
    _ = p0
    _ = v7
    return 728
}

func main() {
    var v1 uint8 = 155
    var v2 float64 = 1.5
    v3 := make(map[uint32]int)
    v4 := []string{(("foo bar" + "") + ""), "", "a"}
    v5 := []int{-998}
    v8 := v3[65535]
    if false {
        v1 = (v1 / ((v1) | 1))
        _ = len(v3)
    }
    var v9 uint8 = v1
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 string = "a"
        v12 := GenS1{F0: 4294967295}
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        switch v8 {
        case 917:
            v2 += 3.14159
        case -151:
            if true {
                var v14 int32 = -785
                _ = v14
            } else {
                if len(v5) > 0 {
                    v5[0] = v8
                }
            }
        }
        v15 := GenS2{F0: -1, F1: 2147483647, F2: 15}
        _ = fi13
        _ = v15
    }
    var v16 bool = false
    var v17 bool = (v16 && v16)
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: -9223372036854775808, F1: int32(fi18), F2: -63}
        var v20 int64 = 0
        if (!(394 < 868)) {
            var v21 float64 = v2
            v16 = (0 > 42)
            _ = v21
        } else {
            _ = len(v5)
            v5 = append(v5, fi18)
            v1 = (v9 & 42)
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
