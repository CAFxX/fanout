package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0() int8 {
    var v5 int16 = 1
    v6 := GenS1{F0: 200}
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint16, p1 uint64) uint16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        if false {
            for fi8 := 0; fi8 < 19; fi8++ {
                fmt.Println(p1, p1)
                p1 -= 2
                v9 := GenS1{F0: 1}
                _ = fi8
                _ = v9
            }
            p0 -= 32767
        } else {
            var v10 int = fi7
            _ = v10
        }
        p1 = p1
        _ = fi7
    }
    v11 := GenS1{F0: -1}
    var v12 float64 = (1e10 + float64(1e10))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 101
}

func main() {
    var v1 float32 = (3.14159 + float32(0.1))
    var v2 int8 = 42
    var v3 uint32 = 65535
    var v4 int = 2147483647
    v13 := GenS1{F0: 42}
    switch v4 {
    case 4294967295:
        v1 = -2.25
    case 1:
        var v14 int32 = 1
        _ = v14
    default:
        v15 := []int{(v4 / (((v4 << ((0) & 63))) | 1)), (v4 >> ((255) & 63)), (^(v4))}
        _ = v15
    }
    if (255 >= 203) {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int = (fi16 >> ((1) & 63))
            if (540 == 42) {
                var v18 uint = 2
                _ = v18
            }
            v19 := []int{v4}
            _ = fi16
            _ = v17
            _ = v19
        }
        var v20 int64 = -1
        v21 := []int{v4}
        _ = v20
        _ = v21
    }
    v2 += (v2 - v2)
    fmt.Println(v3, v4)
    var v22 int64 = -9223372036854775808
    v2 += genFn0()
    v23 := []int{}
    if true {
        v24 := GenS2{F0: 2, F1: true}
        if false {
            if false {
                for fi25 := 0; fi25 < 1; fi25++ {
                    v1 = (float32(-2.25) * 0.0)
                    v22 += (v22 | 9223372036854775807)
                    var v26 float64 = 0.0
                    _ = fi25
                    _ = v26
                }
                v3 = v3
                v1 -= ((((float32(float64(v1)) - 0.0) / v1) + v1) * (v1 * v1))
            } else {
                v1 = (v1 + ((v1 + 1.5) + v1))
                var v27 float64 = (-2.25 + -2.25)
                _ = v27
            }
            fmt.Println(v4)
            fmt.Println(v1, v1)
        } else {
            v23 = append(v23, (v4 + v4))
            v1 = 0.0
            v4 = (v4 - v4)
        }
        _ = v24
    }
    if false {
        var v28 float32 = v1
        var v29 bool = false
        v22 = v22
        _ = v28
        _ = v29
    }
    v30 := GenS2{F0: 32767, F1: ("a" == "foo bar")}
    var v31 int64 = v22
    var v32 bool = ((v4 << ((65535) & 63)) > (v4 + int(v2)))
    if (42 != 1) {
        var v33 int8 = v2
        v23 = append(v23, v4)
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
