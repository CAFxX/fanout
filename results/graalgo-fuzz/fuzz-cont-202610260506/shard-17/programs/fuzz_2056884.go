package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 uint32, p2 uint16) int32 {
    var v6 int32 = 42
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return v6
}
func genFn1() string {
    var v7 uint64 = 0
    fmt.Println(v7)
    v7 -= (v7 % (((v7 >> ((0) & 63))) | 1))
    if false {
        var v8 float32 = 1.5
        v9 := GenS1{F0: 128, F1: 42}
        _ = v8
        _ = v9
    } else {
        var v10 int64 = -830
        _ = v10
    }
    _ = v7
    return "hello"
}

func main() {
    var v1 uint8 = 128
    var v2 float32 = (((0.0 * 3.14159) / 1.5) - 1e10)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v5 float32 = v2
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 float64 = (3.14159 + 1.5)
        var v13 uint16 = 42
        var v14 int8 = -121
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := []int{0, 797, 2147483647, 400}
    fmt.Println(v1, v1)
    v16 := GenS2{F0: ""}
    v17 := GenS1{F0: 500, F1: 42}
    v18 := v4[-910]
    fmt.Println(v18, v18)
    for fi19 := 0; fi19 < 14; fi19++ {
        v18 = v18
        _ = len(v15)
        _ = fi19
    }
    var v20 uint64 = 2
    if false {
        var v21 string = ""
        for fi22 := 0; fi22 < 20; fi22++ {
            for fi23 := 0; fi23 < 4; fi23++ {
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi23
            }
            _ = fi22
        }
        var v24 bool = false
        _ = v21
        _ = v24
    } else {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        switch "hello" {
        case "a":
            var v25 int8 = 42
            _ = v25
        case "":
            if true {
                fmt.Println(v1, v1, v2)
                v26 := GenS1{F0: -1, F1: -71}
                _ = v26
            }
        case "hello":
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 int = (fi27 % (((fi27 * 835)) | 1))
                _ = fi27
                _ = v28
            }
        }
        v5 -= ((v5 / v2) + (v2 * float32(-2.25)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
