package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 string }
func genFn0(p0 uint, p1 int32) int16 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int8 = -112
        var v8 float64 = (float64(1.5) * 1e10)
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(p0, p1)
    case "hello":
        fmt.Println(p1)
    case "":
        p0 -= 65535
    }
    v9 := []int{127, 128, -402}
    _ = p0
    _ = p1
    _ = v9
    return -577
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 uint64 = 4294967295
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    v10 := []int{-2147483648}
    if false {
        v11 := GenS1{F0: 513}
        _ = v11
    } else {
        var v12 float32 = ((0.0 - 0.1) / 1.0)
        _ = v12
    }
    fmt.Println(v3)
    var v13 int8 = -62
    var v14 bool = (v13 >= -128)
    var v15 uint8 = 42
    if true {
        if v14 {
            var v16 int = -171
            var v17 float64 = ((0.0 + -2.25) / 3.14159)
            _ = v16
            _ = v17
        }
        var v18 int8 = v13
        _ = v18
    } else {
        fmt.Println(v15)
        var v19 float64 = float64(0.0)
        _ = v19
    }
    if (!false) {
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 int = (fi20 ^ (fi20 ^ 0))
            v22 := GenS2{F0: 765}
            var v23 uint8 = v15
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        if (true && (((!(32767 <= 128)) && (32768 < 0)) && (v3 <= (v3 >> ((2) & 63))))) {
            if (2 <= 2) {
                v15 = v15
                var v24 string = ("a" + ("foo bar" + "\n\t"))
                _ = v24
            }
            if (!v14) {
                v3 = (v3 - v3)
                if v14 {
                    v25 := GenS1{F0: -327}
                    fmt.Println(v14, v3)
                    _ = v25
                }
                var v26 uint16 = 886
                _ = v26
            } else {
                v27 := GenS1{F0: -2147483648}
                _ = v27
            }
        } else {
            v28 := GenS3{F0: -237}
            _ = v28
        }
    } else {
        var v29 uint8 = v15
        _ = len(v10)
        for fi30 := 0; fi30 < 4; fi30++ {
            for fi31 := 0; fi31 < 8; fi31++ {
                v32 := GenS4{F0: v3}
                fmt.Println(v29)
                v33 := GenS5{F0: uint32(v15), F1: ((((("" + ("" + "foo bar")) + ("\n\t" + ("hello" + ("" + "foo bar")))) + ("a" + "foo bar")) + "") + ("hello" + ("a" + "hello")))}
                _ = fi31
                _ = v32
                _ = v33
            }
            var v34 int32 = 982
            _ = fi30
            _ = v34
        }
        _ = v29
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        fmt.Println(v3, v14, fi35)
        _ = fi35
    }
    v13 = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
