package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint, p1 int64) int16 {
    for fi7 := 0; fi7 < 9; fi7++ {
        switch "" {
        case "a":
            var v8 int32 = 2147483647
            _ = v8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := []int{int(p1), 32768, (fi7 & fi9), fi9}
            v11 := GenS1{F0: false, F1: (("foo bar" + "a") + "hello"), F2: 1}
            var v12 bool = (100 >= 128)
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    if (1.5 >= (1e10 - 0.0)) {
        p0 = p0
        var v13 bool = true
        _ = v13
    }
    p0 += (p0 / ((p0) | 1))
    _ = p0
    _ = p1
    return 1000
}
func genFn1(p0 uint64) uint32 {
    var v14 uint16 = 42
    v14 -= v14
    switch 32768 {
    case 127:
        var v15 uint64 = p0
        _ = v15
    case 0:
        var v16 string = "a"
        _ = v16
    }
    v14 += 255
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 int32 = -1
    v2 := []string{("hello" + "foo bar")}
    v3 := make(map[uint32]int)
    var v4 float64 = (-2.25 * -2.25)
    var v5 uint32 = 4294967295
    var v6 int64 = -41
    var v17 string = "hello"
    for fi18 := 0; fi18 < 17; fi18++ {
        if false {
            var v19 uint32 = v5
            for fi20 := 0; fi20 < 19; fi20++ {
                fmt.Println(fi18)
                v4 = v4
                var v21 uint16 = 255
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        var v22 uint8 = 128
        fmt.Println(v17, v5)
        _ = fi18
        _ = v22
    }
    v5 = (v5 * (v5 - genFn1(0)))
    fmt.Println(v1)
    v1 -= (v1 % ((v1) | 1))
    v6 -= (v6 * v6)
    switch v17 {
    case "a":
        v23 := GenS1{F0: false, F1: ((((v17 + v17) + "a") + (v17 + v17)) + v17), F2: 617}
        _ = v23
    default:
        var v24 int8 = -1
        _ = v24
    }
    var v25 int64 = v6
    for fi26 := 0; fi26 < 5; fi26++ {
        fmt.Println(fi26, v1, v17)
        v6 += v6
        _ = fi26
    }
    v1 += v1
    v27 := GenS2{F0: v17}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
