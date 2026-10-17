package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8) uint64 {
    for fi4 := 0; fi4 < 19; fi4++ {
        p0 -= (^(p0))
        _ = fi4
    }
    v5 := GenS1{F0: ("hello" + "\n\t")}
    v6 := []int{32768, -905, 4294967295, 128}
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint64 = 882
    var v2 int16 = 42
    var v3 uint8 = 1
    v7 := []int{-9223372036854775808, -969, -741}
    var v8 int = 127
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(v1, fi9, v3)
        _ = fi9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        if (!(!true)) {
            switch ("\n\t" + "a") {
            case "":
                var v11 uint16 = 0
                _ = v11
            default:
                var v12 uint8 = (v3 >> ((2) & 63))
                _ = v12
            }
            if false {
                if len(v7) > 0 {
                    v7[0] = fi10
                }
            }
        } else {
            var v13 float64 = 1e10
            fmt.Println(v1, v1, v8)
            _ = v13
        }
        var v14 int16 = v2
        _ = fi10
        _ = v14
    }
    var v15 bool = (!true)
    v8 = 32768
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float32 = (((-2.25 + 0.0) - (-2.25 * 1e10)) + (float32(1.5) / 1.5))
        for fi18 := 0; fi18 < 15; fi18++ {
            if (!true) {
                v7 = append(v7, (-(v8)))
            }
            _ = fi18
        }
        v1 = v1
        _ = fi16
        _ = v17
    }
    v19 := GenS2{F0: ((v15 || v15) && v15)}
    v7 = append(v7, v8)
    v1 = (^(v1))
    var v20 float64 = 3.14159
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v21 uint = 0
    for fi22 := 0; fi22 < 12; fi22++ {
        v20 = (0.1 - v20)
        if v15 {
            v1 = v1
        } else {
            var v23 int = (v8 | (v8 % ((v8) | 1)))
            var v24 int64 = 166
            _ = v23
            _ = v24
        }
        _ = fi22
    }
    switch (v8 * (v8 & -1)) {
    case 822:
        v21 -= (v21 & (v21 / ((1) | 1)))
    default:
        var v25 int64 = 1
        _ = v25
    }
    if len(v7) > 0 {
        v7[0] = v8
    }
    v8 -= 127
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v7
    _ = v8
}
