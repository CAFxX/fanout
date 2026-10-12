package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint }
type GenS2 struct { F0 uint; F1 bool }
func genFn0(p0 int8) uint {
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(fi6)
        var v7 uint8 = 2
        if false {
            for fi8 := 0; fi8 < 16; fi8++ {
                v7 -= 2
                fmt.Println(v7)
                p0 += (p0 << ((0) & 63))
                _ = fi8
            }
        } else {
            fmt.Println(fi6)
            var v9 int16 = 143
            var v10 uint64 = 4294967295
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    p0 += p0
    v11 := []int{2147483647}
    if (!(!true)) {
        var v12 float64 = 0.0
        var v13 uint64 = 2
        v12 = (-2.25 * -2.25)
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v11
    return 392
}
func genFn1(p0 uint, p1 int32, p2 int16) uint32 {
    var v14 int64 = -1
    var v15 uint = (^(p0))
    v16 := GenS1{F0: v14, F1: -755, F2: v15}
    v17 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    v1 := []int{-2147483648, 9223372036854775807, 127, 0}
    var v2 uint32 = 42
    var v3 uint64 = 4294967295
    var v4 int32 = 1
    var v5 bool = false
    var v18 uint16 = 42
    v1 = append(v1, 63)
    switch (("a" + "hello") + "hello") {
    case "foo bar":
        var v19 int = 127
        _ = v19
    default:
        v20 := GenS2{F0: 1, F1: (206 != 127)}
        _ = v20
    }
    var v21 int16 = 1000
    v22 := GenS1{F0: 0, F1: 491, F2: 65535}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v18 += (v18 & 646)
    var v23 string = "foo bar"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v24 float64 = (3.14159 + 1e10)
    for fi25 := 0; fi25 < 7; fi25++ {
        v4 += -201
        _ = fi25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v1 = append(v1, (^(fi26)))
        var v27 int32 = 0
        var v28 bool = (127 == 55)
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v21, v2, v4)
    var v29 int8 = 1
    if v5 {
        fmt.Println(v29)
        fmt.Println(v18, v3, v24)
        v18 = 255
    } else {
        switch (("foo bar" + v23) + "hello") {
        case "a":
            fmt.Println(v24, v2)
        }
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
