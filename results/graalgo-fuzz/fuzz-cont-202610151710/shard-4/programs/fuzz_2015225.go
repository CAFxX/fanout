package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0(p0 uint8, p1 uint8, p2 uint32) uint64 {
    p2 -= (p2 * p2)
    p1 = (p1 + (p1 + p0))
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 string = "a"
    var v2 int = 4294967295
    var v3 int8 = 127
    var v4 uint8 = 238
    var v5 bool = ((1.5 - ((((1.5 / 1.5) - 0.0) / 1.0) / 1.0)) < 0.0)
    v6 := []int{(v2 << ((1) & 63)), (v2 + (-(v2))), -584}
    var v7 uint = 65535
    v7 = 988
    var v8 int32 = 42
    v9 := GenS1{F0: v8}
    if v5 {
        var v10 uint32 = 34
        _ = v10
    } else {
        if v5 {
            if len(v6) > 0 {
                v6[0] = v2
            }
        } else {
            v2 = (v2 - (v2 ^ int(v7)))
        }
        var v11 int16 = 42
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        _ = len(v6)
        var v13 string = v1
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS1{F0: (^(v8))}
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v6 = append(v6, v2)
    switch (v2 | v2) {
    case 4294967295:
        fmt.Println(v7, v4)
    case 0:
        v2 += -842
    }
    var v16 uint32 = 4294967295
    if (((^(v7)) == 255) || false) {
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(v5)
            v18 := GenS2{F0: fi17, F1: 1000}
            for fi19 := 0; fi19 < 1; fi19++ {
                fmt.Println(v1, v7)
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
    }
    if (448 < v8) {
        for fi20 := 0; fi20 < 15; fi20++ {
            for fi21 := 0; fi21 < 15; fi21++ {
                v6 = append(v6, 4294967295)
                var v22 int8 = v3
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        fmt.Println(v1, v16, v8)
        fmt.Println(v5)
    }
    v7 = (v7 + (^(v7)))
    for fi23 := 0; fi23 < 16; fi23++ {
        v3 -= v3
        _ = fi23
    }
    v4 = 2
    v4 = (v4 << ((v7) & 63))
    var v24 bool = ((!false) && true)
    fmt.Println(v24, v8, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
