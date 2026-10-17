package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }
func genFn0(p0 int64, p1 int32) bool {
    p1 = p1
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(p0, p0)
        _ = fi7
    }
    _ = p0
    _ = p1
    return (4294967295 < 18446744073709551615)
}
func genFn1(p0 int16) int8 {
    switch 2147483647 {
    case -336:
        var v8 string = ("foo bar" + "a")
        _ = v8
    case 1:
        v9 := GenS1{F0: 0, F1: -1, F2: 1}
        _ = v9
    case -484:
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(fi10, fi10)
            v11 := []int{-72, (fi10 % (((fi10 >> ((1) & 63))) | 1)), -2147483648, (fi10 % (((fi10 / (((-(fi10))) | 1))) | 1))}
            _ = fi10
            _ = v11
        }
    }
    var v12 uint8 = 1
    _ = p0
    _ = v12
    return 19
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{}
    v3 := make(map[string]bool)
    var v4 uint8 = 0
    var v5 int8 = -61
    var v6 int32 = 800
    v5 -= (^(v5))
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            if true {
                fmt.Println(fi14, v6, fi14)
                v5 = v5
            } else {
                v15 := v3["foo bar"]
                _ = v15
            }
            _ = fi14
        }
        _ = fi13
    }
    var v16 bool = false
    for fi17 := 0; fi17 < 16; fi17++ {
        switch ((("" + "hello") + "foo bar") + "\n\t") {
        case "hello":
            v16 = v16
        default:
            var v18 uint32 = 378
            _ = v18
        }
        _ = fi17
    }
    var v19 int = 4294967295
    var v20 uint32 = 324
    fmt.Println(v6, v5)
    var v21 float64 = 3.14159
    var v22 uint16 = 0
    v4 = (^(v4))
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
