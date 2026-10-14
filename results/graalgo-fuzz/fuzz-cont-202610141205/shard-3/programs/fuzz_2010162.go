package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int32; F2 uint }
func genFn0(p0 uint64, p1 uint64) uint64 {
    if false {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 int32 = -2147483648
            _ = fi5
            _ = v6
        }
    }
    var v7 float64 = -2.25
    var v8 bool = true
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p1 | (p0 << ((255) & 63)))
}

func main() {
    var v1 uint16 = 144
    var v2 string = ("a" + "foo bar")
    var v3 uint64 = 2
    v4 := []int{}
    var v9 int16 = -1
    var v10 int = 32767
    var v11 float64 = (3.14159 + 0.0)
    v10 += (v10 ^ 544)
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 string = v2
        var v14 int8 = -1
        _ = fi12
        _ = v13
        _ = v14
    }
    if len(v4) > 0 {
        v4[0] = 0
    }
    fmt.Println(v10, v1, v10)
    switch (v10 - v10) {
    case -182:
        if (!(166 > 42)) {
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := GenS1{F0: v1}
                _ = fi15
                _ = v16
            }
            v10 = (v10 * (v10 + (v10 - (-(v10)))))
        } else {
            var v17 uint16 = v1
            var v18 int16 = 443
            var v19 float32 = 1e10
            _ = v17
            _ = v18
            _ = v19
        }
    }
    var v20 uint = 4294967295
    var v21 uint = 504
    if len(v4) > 0 {
        v4[0] = int(v21)
    }
    if (!false) {
        if (v9 > 424) {
            _ = len(v4)
        }
        v22 := GenS2{F0: (v10 | 0), F1: 831, F2: uint(v3)}
        _ = v22
    } else {
        var v23 uint64 = (v3 & 4294967295)
        _ = v23
    }
    var v24 float64 = (v11 - v11)
    var v25 string = ((v2 + v2) + v2)
    var v26 float64 = (3.14159 + (v11 - (float64(3.14159) * ((0.1 + (0.0 + (v24 * v24))) + (v11 - v11)))))
    v20 = (v21 + v20)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
