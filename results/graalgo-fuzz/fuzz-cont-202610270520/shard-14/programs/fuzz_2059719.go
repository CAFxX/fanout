package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0() int32 {
    v5 := []int{516, 128}
    v6 := GenS1{F0: 0, F1: false}
    var v7 float32 = 1e10
    v5 = append(v5, 32767)
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 string, p1 uint, p2 int) int16 {
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 bool = true
        _ = fi8
        _ = v9
    }
    p2 = 9223372036854775807
    v10 := []int{p2, (p2 << (((p1 + 137)) & 63)), 32767}
    switch (p0 + ("a" + p0)) {
    case "":
        v11 := GenS1{F0: 205, F1: true}
        _ = v11
    case "\n\t":
        switch "a" {
        case "\n\t":
            p2 = -1
        case "":
            _ = len(v10)
        default:
            var v12 string = ""
            _ = v12
        }
    case "a":
        var v13 int64 = 9223372036854775807
        _ = v13
    default:
        p2 += p2
    }
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -32768
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int = -208
    v4 := make(map[string]uint32)
    _ = len(v2)
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := v1[""]
        _ = fi14
        _ = v15
    }
    var v16 uint16 = 42
    v17 := []int{(v3 & 948)}
    var v18 float64 = float64(1.5)
    _ = len(v4)
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 int64 = 42
        for fi21 := 0; fi21 < 3; fi21++ {
            switch ("a" + "") {
            case "a":
                fmt.Println(fi21, fi19, v20)
            }
            _ = fi21
        }
        for fi22 := 0; fi22 < 3; fi22++ {
            _ = len(v2)
            v3 = fi19
            if true {
                var v23 int32 = 334
                _ = len(v17)
                _ = v23
            } else {
                fmt.Println(v16, fi22)
                v20 = 9223372036854775807
            }
            _ = fi22
        }
        _ = fi19
        _ = v20
    }
    var v24 uint = 323
    v25 := GenS1{F0: (v16 + v16), F1: false}
    var v26 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
