package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 bool, p1 uint8, p2 uint64) int8 {
    fmt.Println(p1, p2, p1)
    var v4 int16 = 264
    for fi5 := 0; fi5 < 10; fi5++ {
        p0 = (3.14159 > ((0.1 * 0.0) + 0.0))
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -44
}

func main() {
    var v1 bool = true
    var v2 uint32 = 0
    v3 := []uint32{}
    v6 := GenS1{F0: 73}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 32768 {
    case -490:
        v1 = (197 > 255)
    case 127:
        var v7 int = -2147483648
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        v1 = (true && v1)
        v1 = v1
        _ = fi8
    }
    var v9 int64 = 293
    _ = len(v3)
    var v10 int16 = 0
    if false {
        var v11 float64 = float64(0.0)
        _ = v11
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v12 := GenS1{F0: 42}
        for fi13 := 0; fi13 < 14; fi13++ {
            fmt.Println(fi13, v2, fi13)
            _ = fi13
        }
        _ = v12
    }
    if v1 {
        if (genFn0(v1, 68, 961) != 1) {
            v10 = (v10 / (((v10 & (-(v10)))) | 1))
            fmt.Println(v10, v9)
        }
        switch "\n\t" {
        case "":
            var v14 uint32 = (v2 * (v2 << ((0) & 63)))
            _ = v14
        }
        var v15 string = ("foo bar" + "hello")
        _ = v15
    } else {
        var v16 int64 = int64(v10)
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
