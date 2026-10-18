package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0() uint16 {
    var v7 uint = 866
    v7 = v7
    if false {
        var v8 int8 = -82
        v7 = 289
        _ = v8
    } else {
        v9 := []int{42}
        var v10 uint64 = 864
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := []int{}
        _ = fi11
        _ = v12
    }
    _ = v7
    return 42
}
func genFn1(p0 int16, p1 string) int8 {
    fmt.Println(p0)
    p1 = "a"
    var v13 float32 = -2.25
    p0 = p0
    _ = p0
    _ = p1
    _ = v13
    return 42
}

func main() {
    var v1 uint8 = 108
    var v2 uint16 = 42
    var v3 float64 = 0.1
    v4 := make(map[string]string)
    v5 := []uint32{151, 69, 0}
    v6 := []uint32{733, 164}
    v14 := GenS1{F0: ((true && true) && false), F1: 0, F2: "\n\t"}
    v2 = genFn0()
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint32 = 2147483647
        _ = fi15
        _ = v16
    }
    if false {
        var v17 uint = 4294967295
        if len(v5) > 0 {
            v5[0] = 65535
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 string = "foo bar"
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        _ = v17
    }
    fmt.Println(v1)
    v3 = float64(1.5)
    v21 := GenS1{F0: false, F1: 128, F2: ""}
    var v22 bool = true
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            if v22 {
                v25 := GenS1{F0: v22, F1: fi24, F2: "\n\t"}
                v3 += v3
                _ = v25
            }
            if (!((0 >= 1) || ((!(42 == genFn1(-1, "hello"))) && v22))) {
                fmt.Println(v22, v22)
                v26 := GenS2{F0: 108, F1: -4}
                fmt.Println(v2, fi23, v1)
                _ = v26
            }
            _ = len(v6)
            var v27 uint = 0
            _ = fi24
            _ = v27
        }
        var v28 float32 = float32(1.5)
        _ = fi23
        _ = v28
    }
    v29 := GenS2{F0: 1, F1: -128}
    v1 += (v1 % ((v1) | 1))
    v2 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
