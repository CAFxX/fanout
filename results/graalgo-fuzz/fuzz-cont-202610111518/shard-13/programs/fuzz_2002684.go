package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int64; F2 uint }
func genFn0(p0 bool) uint16 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 11; fi5++ {
        p0 = p0
        v6 := []int{}
        var v7 float32 = ((0.1 + (0.0 / float32(0.1))) / 1.5)
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    return 255
}

func main() {
    var v1 bool = true
    var v2 int16 = -967
    var v3 int = 420
    v4 := []uint32{903, 4294967295, 42, 322}
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 string = "a"
            fmt.Println(v2)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    fmt.Println(v2)
    fmt.Println(v2, v2)
    fmt.Println(v1, v2, v3)
    v11 := GenS1{F0: ("hello" + "hello")}
    var v12 int16 = v2
    var v13 float64 = ((((1e10 - (float64((3.14159 - (0.1 / -2.25))) + -2.25)) + -2.25) + -2.25) - 0.1)
    fmt.Println(v13)
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v13)
    case "hello":
        v14 := GenS1{F0: "foo bar"}
        _ = v14
    }
    var v15 uint64 = 99
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS1{F0: "a"}
        if (-102 > -1) {
            _ = len(v4)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v18 := GenS2{F0: 2, F1: 1, F2: 218}
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
