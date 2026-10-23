package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 bool }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }

func main() {
    var v1 bool = false
    var v2 uint = 42
    var v3 int64 = 384
    var v4 int64 = v3
    var v5 uint32 = 425
    v6 := []int{-190, 61}
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int16 = 1
            var v10 bool = v1
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(fi7, v1, v3)
        _ = fi7
    }
    v11 := GenS1{F0: 4294967295, F1: (^(v2)), F2: true}
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 float64 = 0.0
            _ = fi13
            _ = v14
        }
        fmt.Println(v3, v2, v4)
        _ = fi12
    }
    var v15 int32 = -406
    v1 = v1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if v1 {
        fmt.Println(v3, v3)
        v16 := GenS1{F0: 18446744073709551615, F1: (v2 ^ v2), F2: (!(v15 > (v15 + v15)))}
        fmt.Println(v1, v3)
        _ = v16
    } else {
        v2 = (v2 % ((4294967295) | 1))
        v17 := GenS2{F0: uint16(v2), F1: (255 >= 42), F2: 4294967295}
        _ = v17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 uint8 = 226
        fmt.Println(v2)
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 0
    switch "hello" {
    case "a":
        if v1 {
            _ = len(v6)
            for fi21 := 0; fi21 < 9; fi21++ {
                for fi22 := 0; fi22 < 9; fi22++ {
                    fmt.Println(v5, v4, v3)
                    var v23 uint64 = 1
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
        } else {
            fmt.Println(v3, v1, v20)
            v24 := GenS2{F0: 0, F1: v1, F2: 0}
            fmt.Println(v3)
            _ = v24
        }
    }
    v25 := GenS2{F0: 812, F1: (42 == (v15 ^ v15)), F2: v2}
    if (v2 > 943) {
        var v26 uint = v2
        var v27 uint8 = 255
        var v28 int8 = -78
        _ = v26
        _ = v27
        _ = v28
    }
    fmt.Println(v4, v20)
    var v29 uint64 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v20
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
