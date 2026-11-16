package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int, p1 int64, p2 int8) int16 {
    switch p0 {
    case 574:
        var v5 float32 = float32(-2.25)
        _ = v5
    case 42:
        if ("hello" != "a") {
            v6 := []int{127}
            _ = v6
        }
    default:
        v7 := []int{-765, 2147483647, 2147483647}
        _ = v7
    }
    if true {
        if false {
            v8 := []int{}
            p2 -= (^(p2))
            if len(v8) > 0 {
                v8[0] = 42
            }
            _ = v8
        }
        var v9 int64 = (p1 - (p1 + (p1 ^ (p1 / (((-(p1))) | 1)))))
        var v10 float32 = (((1e10 - 1e10) * (3.14159 + 0.0)) - 3.14159)
        _ = v10
        _ = v9
    }
    fmt.Println(p2, p1, p2)
    for fi11 := 0; fi11 < 8; fi11++ {
        if false {
            v12 := []int{(p0 + (fi11 | 433)), (fi11 ^ p0)}
            _ = v12
        }
        var v13 int32 = 871
        _ = fi11
        _ = v13
    }
    switch p0 {
    case 32768:
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS1{F0: -2147483648, F1: 32767, F2: 140}
            var v16 int64 = p1
            _ = fi14
            _ = v15
            _ = v16
        }
    case 1:
        if false {
            v17 := GenS2{F0: 42, F1: (p2 + p2)}
            fmt.Println(p2, p2, p2)
            _ = v17
        }
    case 42:
        fmt.Println(p1, p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() bool {
    if true {
        var v18 bool = (true || true)
        v19 := []int{}
        v20 := GenS1{F0: 0, F1: 32767, F2: 127}
        _ = v18
        _ = v19
        _ = v20
    } else {
        switch "hello" {
        case "foo bar":
            v21 := []int{}
            _ = v21
        case "":
            var v22 string = "foo bar"
            _ = v22
        case "a":
            for fi23 := 0; fi23 < 13; fi23++ {
                var v24 bool = true
                _ = fi23
                _ = v24
            }
        }
        var v25 int16 = 1000
        for fi26 := 0; fi26 < 20; fi26++ {
            v27 := []int{(^(fi26)), (-(fi26)), (fi26 % (((fi26 * fi26)) | 1))}
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    var v28 int32 = -1
    _ = v28
    return (0 == 65535)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = (0.1 + 0.0)
    v3 := make(map[string]bool)
    var v4 string = ""
    _ = len(v1)
    v2 -= (v2 * 3.14159)
    var v29 int16 = 1000
    v29 = (-(v29))
    for fi30 := 0; fi30 < 1; fi30++ {
        var v31 float64 = 0.1
        _ = fi30
        _ = v31
    }
    var v32 bool = (-2147483648 < 2147483647)
    var v33 int8 = 15
    v34 := GenS3{F0: 1}
    v35 := []int{}
    for fi36 := 0; fi36 < 15; fi36++ {
        v37 := GenS3{F0: 1}
        _ = fi36
        _ = v37
    }
    if v32 {
        fmt.Println(v33, v29)
    } else {
        if (0 == 385) {
            fmt.Println(v4, v4)
        }
        switch -1 {
        case 32768:
            v38 := GenS1{F0: 0, F1: v29, F2: 128}
            _ = v38
        }
        var v39 int64 = 0
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
