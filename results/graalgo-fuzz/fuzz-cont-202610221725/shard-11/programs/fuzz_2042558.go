package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint = 223
    var v2 uint16 = 65535
    var v3 uint32 = 852
    var v4 float32 = (-2.25 / (float32(float64(0.0)) + float32(-2.25)))
    if false {
        var v5 int32 = 1
        fmt.Println(v1, v1, v1)
        _ = v5
    }
    v6 := GenS1{F0: 1, F1: 59}
    v3 += 343
    switch 280 {
    case 985:
        var v7 int8 = -128
        _ = v7
    }
    fmt.Println(v2, v1)
    v2 += 42
    var v8 float32 = (v4 * (-2.25 / 1.5))
    fmt.Println(v4, v1, v3)
    v9 := []int{32768, -969, 2147483647}
    switch int(v3) {
    case 128:
        var v10 int8 = 42
        _ = v10
    case 332:
        switch ("foo bar" + "foo bar") {
        case "":
            v4 -= v8
        case "hello":
            v4 = (v4 * (v8 * v4))
        case "a":
            for fi11 := 0; fi11 < 14; fi11++ {
                v4 = 1e10
                var v12 string = "a"
                if true {
                    v13 := GenS1{F0: 2, F1: 127}
                    v12 = v12
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
        }
    case -687:
        switch "hello" {
        case "hello":
            v14 := GenS1{F0: uint64(v1), F1: -1}
            _ = v14
        }
    }
    if false {
        var v15 int16 = 42
        v16 := GenS2{F0: true}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS2{F0: false}
        var v19 float32 = v4
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 uint = 1
        _ = fi20
        _ = v21
    }
    _ = len(v9)
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 uint16 = (v2 ^ (v2 | 2))
            fmt.Println(v2)
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
