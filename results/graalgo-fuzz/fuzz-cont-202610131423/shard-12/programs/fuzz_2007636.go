package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16; F1 int }

func main() {
    var v1 bool = (32767 <= 1)
    v2 := []int{-1, 42, 32767}
    var v3 bool = ((!true) && (42 == 123))
    var v4 float64 = (0.0 + -2.25)
    v5 := make(map[string]string)
    switch "a" {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        if (((2147483647 == 42) || v1) || v3) {
            for fi6 := 0; fi6 < 8; fi6++ {
                var v7 uint64 = 18446744073709551615
                _ = fi6
                _ = v7
            }
        }
    }
    v4 = (v4 - ((v4 / v4) + float64(((1.5 - float32((0.0 - (v4 - v4)))) + (1e10 - 3.14159)))))
    v1 = (100 < 1)
    v4 = (v4 + v4)
    var v8 float64 = float64(1e10)
    var v9 uint8 = 0
    if v3 {
        if (1 == 0) {
            fmt.Println(v9)
        } else {
            v10 := GenS1{F0: 886, F1: 55}
            switch "hello" {
            case "hello":
                var v11 int64 = 42
                _ = v11
            }
            if len(v2) > 0 {
                v2[0] = -901
            }
            _ = v10
        }
    } else {
        v5[("\n\t" + ((("foo bar" + "") + "a") + "a"))] = "\n\t"
        v12 := GenS1{F0: 42, F1: 42}
        var v13 uint32 = 845
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: -1, F1: 0}
        var v16 int64 = 1
        _ = fi14
        _ = v15
        _ = v16
    }
    if v1 {
        v17 := GenS2{F0: 430, F1: 42}
        v18 := GenS2{F0: 255, F1: int8(v9)}
        _ = v17
        _ = v18
    } else {
        if (608 < 0) {
            var v19 int = 4294967295
            var v20 uint8 = v9
            fmt.Println(v4, v3, v8)
            _ = v19
            _ = v20
        } else {
            _ = len(v2)
        }
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 uint8 = (v9 % (((v9 % ((v9) | 1))) | 1))
            _ = fi22
            _ = v23
        }
        if (((float64(0.0) * (0.0 + (1e10 - (float64(-2.25) - 1.5)))) + -2.25) == -2.25) {
            v3 = true
            v24 := GenS3{F0: 42, F1: fi21}
            v25 := GenS3{F0: 32767, F1: fi21}
            _ = v24
            _ = v25
        }
        v1 = v3
        _ = fi21
    }
    var v26 uint16 = 210
    var v27 int64 = -1
    var v28 int = 32768
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 int = fi29
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
