package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 string }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool) int {
    var v5 uint32 = 4294967295
    fmt.Println(v5, p0)
    var v6 uint8 = 194
    v6 += (v6 / ((v6) | 1))
    var v7 int16 = 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 230
}
func genFn1(p0 uint16) int16 {
    var v8 uint = 361
    p0 = uint16(v8)
    _ = p0
    _ = v8
    return 32767
}

func main() {
    var v1 int64 = -1
    var v2 int64 = (v1 + (-(v1)))
    var v3 uint32 = 2147483647
    var v4 uint64 = 0
    v9 := []int{127}
    var v10 int16 = 163
    v3 -= (v3 & v3)
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: fi11, F1: 0, F2: (("a" + "") + ("" + "\n\t"))}
        v4 += 927
        v9 = append(v9, -779)
        _ = fi11
        _ = v12
    }
    v9 = append(v9, 628)
    for fi13 := 0; fi13 < 7; fi13++ {
        if true {
            var v14 uint = 0
            v3 = (v3 | (v3 % ((294) | 1)))
            v4 += 1
            _ = v14
        } else {
            fmt.Println(v1, v2, v3)
        }
        var v15 uint32 = 2147483647
        _ = fi13
        _ = v15
    }
    if (3.14159 == 0.0) {
        v16 := GenS2{F0: 4294967295}
        _ = v16
    } else {
        v17 := GenS1{F0: 1, F1: 838, F2: "foo bar"}
        _ = v17
    }
    if len(v9) > 0 {
        v9[0] = 42
    }
    fmt.Println(v4, v1, v4)
    for fi18 := 0; fi18 < 6; fi18++ {
        _ = len(v9)
        v10 = v10
        _ = fi18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v9 = append(v9, fi19)
        v20 := GenS2{F0: v3}
        var v21 int64 = (v2 % ((42) | 1))
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v1, v1)
    var v22 uint32 = (v3 / ((639) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
