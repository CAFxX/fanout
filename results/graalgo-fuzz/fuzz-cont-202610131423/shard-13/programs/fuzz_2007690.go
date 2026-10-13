package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 10; fi6++ {
        if false {
            var v7 float32 = (0.1 + 3.14159)
            _ = v7
        } else {
            var v8 int32 = 1
            if true {
                var v9 uint32 = 1
                _ = v9
            }
            v10 := GenS1{F0: 313}
            _ = v10
            _ = v8
        }
        fmt.Println(fi6)
        _ = fi6
    }
    if (!(42 >= 0)) {
        var v11 float32 = 0.1
        _ = v11
    } else {
        var v12 int = 9223372036854775807
        v12 = (v12 ^ v12)
        _ = v12
    }
    var v13 int32 = 1
    _ = v13
    return 330
}
func genFn1(p0 int16, p1 int32) int32 {
    p1 -= (p1 | p1)
    var v14 uint = 42
    var v15 uint = (v14 - 255)
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return p1
}

func main() {
    var v1 int16 = 1
    var v2 uint64 = 952
    v3 := []uint32{65535}
    v4 := []uint32{}
    var v5 uint16 = 42
    if ((v2 | (v2 * v2)) <= v2) {
        for fi16 := 0; fi16 < 7; fi16++ {
            fmt.Println(v2, v5)
            _ = fi16
        }
    } else {
        var v17 uint16 = (v5 << ((255) & 63))
        _ = v17
    }
    v1 -= -543
    for fi18 := 0; fi18 < 8; fi18++ {
        _ = len(v3)
        v2 = 42
        var v19 uint32 = genFn0()
        _ = fi18
        _ = v19
    }
    fmt.Println(v2, v5, v1)
    var v20 string = "foo bar"
    v5 += (^(v5))
    v21 := GenS2{F0: -2147483648}
    if true {
        v5 -= 316
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
