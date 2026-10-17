package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint16, p1 bool, p2 string) int64 {
    p1 = p1
    p2 = (((p2 + p2) + "") + "foo bar")
    if false {
        var v7 uint = 2
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{(fi8 / ((fi8) | 1)), fi8}
            v7 -= v7
            _ = fi8
            _ = v9
        }
        v7 += v7
        _ = v7
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := []int{fi10, 42, (^(fi10)), (fi10 >> ((0) & 63))}
        var v12 float64 = (1.5 - (1e10 - ((0.0 + float64(-2.25)) - 0.0)))
        for fi13 := 0; fi13 < 16; fi13++ {
            p2 = "a"
            p1 = p1
            fmt.Println(p0)
            _ = fi13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := []int{187}
        fmt.Println(p2, p2, p1)
        var v16 float64 = 1.5
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1() int64 {
    v17 := []int{127, -182}
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(fi18)
        _ = fi18
    }
    var v19 uint64 = 179
    _ = v17
    _ = v19
    return 268
}

func main() {
    var v1 uint8 = 0
    var v2 uint8 = 166
    v3 := make(map[string]string)
    var v4 int = 42
    var v5 uint8 = 127
    var v6 uint64 = 2
    v20 := GenS1{F0: (-2147483648 == -932), F1: v4}
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v1)
        var v22 int = fi21
        v5 = (v2 >> ((107) & 63))
        _ = fi21
        _ = v22
    }
    v2 = v5
    var v23 float64 = 1e10
    if (0 != 42) {
        fmt.Println(v4, v2)
        v6 = (v6 + 18446744073709551615)
    }
    var v24 int64 = -369
    v25 := GenS1{F0: (!true), F1: (v4 - 9223372036854775807)}
    v26 := GenS2{F0: v4}
    v27 := []int{(-(v4))}
    v28 := GenS3{F0: "hello", F1: false, F2: 42}
    switch (v4 % ((-9223372036854775808) | 1)) {
    case -385:
        v29 := GenS4{F0: uint32(v5), F1: -128}
        _ = v29
    }
    fmt.Println(v4, v6, v4)
    v4 -= (v4 % (((v4 - (v4 & 32767))) | 1))
    for fi30 := 0; fi30 < 1; fi30++ {
        v23 = ((v23 + (v23 / 0.1)) * (((v23 * v23) * float64(1.5)) - float64(-2.25)))
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 bool = (!(true || true))
            var v33 int64 = (v24 ^ 0)
            v24 -= 0
            _ = fi31
            _ = v32
            _ = v33
        }
        var v34 uint32 = 4294967295
        _ = fi30
        _ = v34
    }
    v1 = v5
    for fi35 := 0; fi35 < 2; fi35++ {
        fmt.Println(v24, v6)
        var v36 int32 = 541
        fmt.Println(v24, v5)
        _ = fi35
        _ = v36
    }
    fmt.Println(v1)
    var v37 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
