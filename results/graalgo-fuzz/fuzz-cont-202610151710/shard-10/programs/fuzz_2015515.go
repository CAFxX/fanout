package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    var v5 int16 = 1000
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(p0, p0)
        var v7 uint16 = 32767
        _ = fi6
        _ = v7
    }
    switch 32767 {
    case -602:
        v8 := GenS1{F0: 0}
        _ = v8
    case 32768:
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int64 = 243
            if (15 >= 127) {
                var v11 int8 = 86
                if false {
                    fmt.Println(p0, v11)
                    var v12 uint16 = 660
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 3; fi13++ {
                        var v14 uint32 = 951
                        _ = fi13
                        _ = v14
                    }
                    v15 := GenS2{F0: (-(fi9))}
                    _ = v15
                }
                v16 := GenS1{F0: v10}
                _ = v11
                _ = v16
            } else {
                for fi17 := 0; fi17 < 17; fi17++ {
                    v18 := []int{9223372036854775807, fi17}
                    if (v10 != 0) {
                        var v19 uint32 = 42
                        _ = v19
                    }
                    _ = fi17
                    _ = v18
                }
            }
            v10 = v10
            _ = fi9
            _ = v10
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        fmt.Println(p0)
        var v21 string = "foo bar"
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = v5
    return 128
}

func main() {
    var v1 int16 = 851
    var v2 uint32 = 2147483648
    var v3 bool = (1 > 100)
    var v4 uint64 = 42
    var v22 float64 = (0.1 + 0.1)
    var v23 int64 = -9223372036854775808
    switch 32768 {
    case 2147483647:
        v24 := []int{127, 32768, 127}
        _ = v24
    default:
        var v25 uint8 = 91
        _ = v25
    }
    v1 += -32768
    v22 = v22
    v26 := []int{-79, 127, 32767}
    v23 = v23
    var v27 float64 = v22
    if (v23 < -1) {
        var v28 uint = 0
        v27 -= v27
        _ = v28
    }
    v29 := GenS1{F0: -880}
    var v30 int8 = 100
    if true {
        if (v3 || (!(!v3))) {
            v26 = append(v26, 0)
            var v31 float64 = (3.14159 * v27)
            v23 = 778
            _ = v31
        } else {
            if len(v26) > 0 {
                v26[0] = 32768
            }
        }
        v32 := GenS1{F0: v23}
        _ = v32
    }
    if (!(42 >= 127)) {
        _ = len(v26)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
