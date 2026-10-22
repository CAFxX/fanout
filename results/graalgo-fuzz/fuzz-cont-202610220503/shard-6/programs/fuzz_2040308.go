package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int }
func genFn0(p0 int64, p1 uint8, p2 int8) uint16 {
    fmt.Println(p1, p2, p1)
    if (false && true) {
        var v7 int8 = 0
        _ = v7
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := []int{fi8}
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int8 = 64
        v11 = (p2 ^ 1)
        p1 = (p1 << ((550) & 63))
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 168
}
func genFn1(p0 string) uint8 {
    p0 = (p0 + (p0 + (((("" + "hello") + p0) + (p0 + "foo bar")) + p0)))
    v14 := GenS1{F0: 1}
    var v15 float32 = float32((0.1 + 1.5))
    var v16 uint16 = 1
    v17 := []int{9223372036854775807}
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 float64 = 1e10
    var v2 int = 127
    var v3 uint32 = 42
    var v4 uint32 = (v3 << ((18446744073709551615) & 63))
    v5 := []string{"\n\t"}
    v6 := make(map[int]string)
    var v18 int16 = -1
    fmt.Println(v18, v1)
    fmt.Println(v1, v1)
    var v19 int = (v2 - 0)
    v20 := GenS1{F0: 843}
    var v21 uint16 = 32767
    if true {
        var v22 bool = false
        var v23 uint = 516
        var v24 int16 = (^(v18))
        _ = v22
        _ = v23
        _ = v24
    } else {
        var v25 uint16 = v21
        _ = v25
    }
    switch "" {
    case "foo bar":
        v1 = 0.1
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 float32 = 0.1
        v27 = v27
        v18 += v18
        _ = fi26
        _ = v27
    }
    if false {
        v1 += float64(0.1)
    } else {
        v28 := GenS2{F0: -9223372036854775808, F1: 2147483647}
        _ = v28
    }
    var v29 int = v19
    v2 = (v19 | v19)
    _ = len(v5)
    v18 = v18
    var v30 int = v29
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
