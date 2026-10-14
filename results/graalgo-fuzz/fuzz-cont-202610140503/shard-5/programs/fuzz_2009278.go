package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }

func main() {
    var v1 bool = (1 < -758)
    var v2 float64 = 1e10
    var v3 uint8 = 128
    v4 := make(map[int]uint32)
    v4[32768] = 42
    var v5 uint64 = 260
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint32 = 2
        var v8 bool = (v1 && v1)
        v2 += (0.0 + -2.25)
        _ = fi6
        _ = v7
        _ = v8
    }
    v2 = float64(float32(v2))
    v1 = false
    var v9 int64 = -1
    var v10 float32 = -2.25
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(v2, v1)
        v12 := GenS1{F0: (^(v3)), F1: "foo bar", F2: 42}
        v2 += (v2 + v2)
        _ = fi11
        _ = v12
    }
    v13 := []int{-388, 944, 42, 4294967295}
    var v14 int = -309
    v10 += (float32(v2) * v10)
    if true {
        v4[(v14 * (v14 << ((4294967295) & 63)))] = 590
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := GenS1{F0: uint8(v5), F1: "a", F2: 42}
            v17 := GenS2{F0: 1, F1: 1000, F2: 0}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int8 = 0
        _ = v18
    }
    switch v14 {
    case 4294967295:
        v10 = (((float32(v2) - v10) * v10) * -2.25)
    }
    v19 := GenS3{F0: uint(v3), F1: 65535, F2: v9}
    if ("" == (((("foo bar" + "hello") + "foo bar") + "\n\t") + "a")) {
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int32 = 116
                var v23 uint8 = (v3 - v3)
                var v24 float64 = 0.0
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            switch 32767 {
            case 32768:
                var v25 bool = v1
                _ = v25
            case 2147483647:
                fmt.Println(v9, fi20)
            case 128:
                for fi26 := 0; fi26 < 17; fi26++ {
                    delete(v4, fi26)
                    _ = len(v13)
                    v9 = v9
                    _ = fi26
                }
            }
            v2 -= (v2 / 1.5)
            _ = fi20
        }
    } else {
        for fi27 := 0; fi27 < 9; fi27++ {
            if v1 {
                for fi28 := 0; fi28 < 2; fi28++ {
                    v29 := v4[v14]
                    var v30 string = ""
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                switch (v14 % ((v14) | 1)) {
                case 32767:
                    v31 := GenS3{F0: 0, F1: 1, F2: v9}
                    _ = v31
                case 432:
                    v3 = (v3 << ((699) & 63))
                }
                _ = len(v13)
            }
            _ = fi27
        }
    }
    v9 -= (v9 + (v9 / ((9223372036854775807) | 1)))
    var v32 int = v14
    _ = len(v13)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
