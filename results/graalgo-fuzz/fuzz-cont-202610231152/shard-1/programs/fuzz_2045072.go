package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int64; F1 uint }
func genFn0() uint16 {
    var v4 uint8 = 128
    v4 += (^(v4))
    _ = v4
    return 42
}
func genFn1(p0 int64, p1 int16) uint {
    p0 -= (p0 << ((322) & 63))
    var v5 string = ("hello" + "hello")
    var v6 int64 = (p0 >> ((18446744073709551615) & 63))
    fmt.Println(p0, p0, p1)
    var v7 float64 = ((1e10 * 0.0) + (0.1 * 1e10))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 bool = true
    v8 := []int{128, 177, 248, -56}
    fmt.Println(v3)
    v9 := GenS1{F0: ((42 < 0) || v3), F1: 433, F2: 358}
    if v3 {
        if v3 {
            v10 := GenS1{F0: v3, F1: 99, F2: 0}
            _ = v10
        } else {
            v3 = v3
            var v11 float32 = (1e10 * (1.5 * 1.5))
            _ = v11
        }
        fmt.Println(v3)
        var v12 int64 = 1
        _ = v12
    } else {
        fmt.Println(v3)
        v13 := GenS2{F0: 961, F1: 42}
        for fi14 := 0; fi14 < 5; fi14++ {
            v3 = (1.5 != (1e10 - (0.0 / 1.5)))
            var v15 int16 = -234
            var v16 string = "hello"
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    if len(v8) > 0 {
        v8[0] = 32767
    }
    var v17 string = ("\n\t" + ("a" + ("\n\t" + "foo bar")))
    var v18 float32 = -2.25
    var v19 uint8 = 0
    switch ("hello" + "") {
    case "foo bar":
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 int = (fi20 >> ((191) & 63))
            _ = fi20
            _ = v21
        }
    case "":
        var v22 bool = (653 != 165)
        _ = v22
    case "a":
        fmt.Println(v18, v17, v3)
    default:
        v23 := GenS3{F0: 0, F1: 702}
        _ = v23
    }
    if v3 {
        for fi24 := 0; fi24 < 5; fi24++ {
            v3 = (32767 < 156)
            _ = fi24
        }
    } else {
        v25 := GenS3{F0: -1, F1: 0}
        v19 = v19
        switch (v17 + "a") {
        case "hello":
            var v26 float64 = 1.5
            _ = v26
        }
        _ = v25
    }
    var v27 int8 = -128
    v3 = v3
    fmt.Println(v19, v19, v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v8
    _ = v9
}
