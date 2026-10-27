package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }
func genFn0() int8 {
    if false {
        switch "foo bar" {
        case "foo bar":
            var v4 uint32 = 961
            _ = v4
        case "":
            var v5 int16 = -665
            _ = v5
        case "hello":
            var v6 uint8 = 0
            _ = v6
        }
    }
    var v7 float64 = 0.1
    var v8 string = "a"
    fmt.Println(v7, v7)
    fmt.Println(v7, v8)
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint64) uint64 {
    var v9 int8 = genFn0()
    v10 := []int{9223372036854775807, 32768, 9223372036854775807, -349}
    if false {
        for fi11 := 0; fi11 < 20; fi11++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi11
        }
    } else {
        var v12 int8 = (-(v9))
        var v13 uint32 = 1
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v9
    return p0
}

func main() {
    v1 := []int{-384, 0, 9223372036854775807}
    var v2 uint64 = 464
    var v3 float64 = float64(1e10)
    fmt.Println(v3, v2)
    var v14 bool = false
    fmt.Println(v14, v2)
    var v15 float64 = v3
    v2 -= (v2 & 152)
    v15 = v3
    fmt.Println(v15)
    v14 = v14
    var v16 int32 = 0
    v16 += (v16 | (-(v16)))
    if false {
        var v17 int32 = (v16 / (((v16 * (v16 + v16))) | 1))
        _ = v17
    }
    fmt.Println(v14, v15, v14)
    fmt.Println(v15)
    for fi18 := 0; fi18 < 4; fi18++ {
        v3 = v15
        for fi19 := 0; fi19 < 7; fi19++ {
            fmt.Println(v16, fi19, v2)
            var v20 uint16 = 0
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = 0
    }
    v21 := GenS1{F0: 4294967295, F1: -12, F2: -2147483648}
    v3 += (((v3 / 1.5) * v15) - v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
}
