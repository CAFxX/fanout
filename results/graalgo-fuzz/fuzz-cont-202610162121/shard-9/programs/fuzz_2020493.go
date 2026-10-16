package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64) uint8 {
    var v5 uint = 4294967295
    v6 := GenS1{F0: 128, F1: "a"}
    fmt.Println(p0)
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            if (false && (!true)) {
                v5 -= (v5 << ((v5) & 63))
            }
            _ = fi8
        }
        var v9 int32 = 42
        v10 := []int{}
        _ = fi7
        _ = v10
        _ = v9
    }
    var v11 uint8 = 2
    _ = p0
    _ = v11
    _ = v5
    _ = v6
    return v11
}

func main() {
    var v1 int = 435
    v2 := []int{127}
    var v3 int64 = 270
    var v4 int64 = v3
    _ = len(v2)
    switch "hello" {
    case "foo bar":
        if true {
            switch "\n\t" {
            case "":
                if true {
                    v4 -= v4
                    v12 := GenS2{F0: 1000}
                    _ = v12
                }
            default:
                v3 -= (v4 & 864)
            }
            fmt.Println(v1, v3, v1)
        } else {
            switch "\n\t" {
            case "hello":
                v4 += v4
            case "":
                var v13 uint = 4294967295
                _ = v13
            default:
                _ = len(v2)
            }
            var v14 int16 = -32768
            var v15 uint = 341
            _ = v14
            _ = v15
        }
    }
    var v16 int32 = -137
    var v17 int16 = 1000
    var v18 int16 = 855
    var v19 uint8 = 255
    for fi20 := 0; fi20 < 18; fi20++ {
        v17 = v17
        if len(v2) > 0 {
            v2[0] = 0
        }
        _ = fi20
    }
    if len(v2) > 0 {
        v2[0] = v1
    }
    v16 = v16
    if len(v2) > 0 {
        v2[0] = (v1 >> ((0) & 63))
    }
    if true {
        var v21 int = (v1 | v1)
        if ((v16 + -2147483648) != (v16 + v16)) {
            v4 = v4
            var v22 uint64 = 42
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 int64 = v4
                if ((^(v21)) > (-(v21))) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v25 bool = (255 > 42)
                    for fi26 := 0; fi26 < 9; fi26++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v21 = fi26
                        if (!(!(!(!true)))) {
                            var v27 int16 = (v17 & v18)
                            _ = len(v2)
                            _ = v27
                        }
                        _ = fi26
                    }
                    _ = v25
                }
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        _ = v21
    } else {
        var v28 float64 = ((0.0 + 1.5) + 1.5)
        fmt.Println(v19, v4, v3)
        if false {
            v4 = v3
            var v29 int8 = -128
            _ = v29
        }
        _ = v28
    }
    v30 := GenS1{F0: genFn0(0), F1: "foo bar"}
    var v31 int64 = (v3 - (v3 & (v3 << ((0) & 63))))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
