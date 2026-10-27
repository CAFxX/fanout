package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
func genFn0(p0 uint16, p1 string, p2 uint) int {
    if (1 >= 0) {
        p0 -= (^(p0))
        v5 := []int{127, 9223372036854775807, 454, 127}
        v6 := GenS1{F0: p2, F1: 4294967295}
        _ = v5
        _ = v6
    }
    if (("a" + "") <= p1) {
        v7 := GenS1{F0: 2, F1: (p2 & p2)}
        p0 += 2
        p2 = p2
        _ = v7
    }
    if (!false) {
        switch -378 {
        case 32767:
            var v8 float64 = float64(0.0)
            _ = v8
        case 267:
            fmt.Println(p2, p0)
        case 0:
            fmt.Println(p1, p1)
        }
        v9 := GenS1{F0: (^(p2)), F1: p2}
        v10 := GenS1{F0: uint(p0), F1: p2}
        _ = v10
        _ = v9
    } else {
        var v11 int = 32767
        var v12 float32 = 0.0
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 902
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    v3 := make(map[string]uint32)
    var v4 bool = ((1000 == -501) || (42 < 101))
    var v13 uint16 = 0
    fmt.Println(v4, v4)
    v13 = v13
    switch (("\n\t" + "hello") + "hello") {
    case "\n\t":
        var v14 uint32 = 65535
        _ = v14
    case "hello":
        fmt.Println(v13, v13)
    default:
        var v15 int = 42
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v13 += v13
            _ = fi17
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4)
        _ = fi16
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    if (4294967295 <= 4294967295) {
        if false {
            v4 = v4
            v18 := GenS1{F0: 65535, F1: 1}
            fmt.Println(v13, v4)
            _ = v18
        }
    }
    fmt.Println(v13)
    fmt.Println(v13, v13)
    v4 = true
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS1{F0: 857, F1: 65535}
        _ = fi19
        _ = v20
    }
    var v21 uint16 = 175
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
